package services

import (
	"errors"

	"github.com/Uttamnath64/quixzap/internal/app/common"
	"github.com/Uttamnath64/quixzap/internal/app/common/types"
	"github.com/Uttamnath64/quixzap/internal/app/models"
	"github.com/Uttamnath64/quixzap/internal/app/repositories"
	"github.com/Uttamnath64/quixzap/internal/app/storage"
	"github.com/Uttamnath64/quixzap/internal/app/templates"
	"github.com/Uttamnath64/quixzap/internal/app/utils"
	"github.com/Uttamnath64/quixzap/internal/app/utils/requests"
	"github.com/Uttamnath64/quixzap/internal/app/utils/responses"
	"gorm.io/gorm"
)

type Auth struct {
	container    *storage.Container
	memberRepo   repositories.MemberRepository
	authRepo     repositories.AuthRepository
	avatarRepo   repositories.AvatarRepository
	jwtUtil      *utils.JWT
	otpService   OTPService
	emailService EmailService
}

func NewAuth(container *storage.Container) *Auth {
	authRepo := repositories.NewAuth(container)
	return &Auth{
		container:    container,
		memberRepo:   repositories.NewMember(container),
		authRepo:     authRepo,
		avatarRepo:   repositories.NewAvatar(container),
		jwtUtil:      utils.NewJWT(container, authRepo),
		otpService:   NewOTP(container.Redis, 300),
		emailService: NewEmail(container),
	}
}

func (service *Auth) Login(rctx *requests.RequestContext, payload requests.Login, deviceInfo string, ip string) responses.ServiceResponse {
	var member models.Member

	// Check member
	if err := service.memberRepo.GetMemberByUsernameOrEmail(rctx, payload.UsernameEmail, payload.UsernameEmail, &member); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Invalid username or email address.", err)
		}

		service.container.Logger.Error("auth.service.login-GetMemberByUsernameOrEmail", "error", err.Error(), "payload", payload)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	// Validate password
	if err := Validate.VerifyPassword(member.Password, payload.Password); err != nil {
		return responses.ErrorResponse(common.StatusValidationError, "Incorrect password. Please try again.", err)
	}
	userType := types.UserTypeMember
	if member.Role == types.MemberRoleSupport {
		userType = types.UserTypeSupportMember
	}

	// Create Token
	accessToken, refreshToken, err := service.jwtUtil.GenerateToken(rctx, member.ID, userType, deviceInfo, ip)
	if err != nil {
		service.container.Logger.Error("auth.service.login-GenerateToken", "error", err.Error(), "userId", member.ID, "userType", userType, "deviceInfo", deviceInfo, "ip", ip)
		return responses.ErrorResponse(common.StatusServerError, "Failed to generate tokens. Please try again later.", err)
	}

	// Response
	return responses.SuccessResponse("Welcome back! 👋 You’re logged in", responses.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func (service *Auth) Register(rctx *requests.RequestContext, payload requests.Register, deviceInfo string, ip string) responses.ServiceResponse {
	var password string

	// Check username
	if err := service.memberRepo.UsernameExists(rctx, payload.Username); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			service.container.Logger.Error("auth.service.register-UsernameExists", "error", err.Error(), "username", payload.Username)
			return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
		}
	} else {
		return responses.ErrorResponse(common.StatusValidationError, "This username is already taken. Please choose another one.", errors.New("username already exists"))
	}

	// Check email
	if err := service.memberRepo.EmailExists(rctx, payload.Email); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			service.container.Logger.Error("auth.service.register-EmailExists", "error", err.Error(), "email", payload.Email)
			return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
		}
	} else {
		return responses.ErrorResponse(common.StatusValidationError, "This email address is already registered. Try logging in or use a different one.", errors.New("email already exists"))
	}

	// Verify avatar
	if err := service.avatarRepo.AvatarByTypeExists(rctx, payload.AvatarId, types.AvatarTypeUser); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusValidationError, "Selected avatar is invalid or does not exist.", errors.New("avatar not found"))
		}
		service.container.Logger.Error("auth.service.register-AvatarByTypeExists", "error", err.Error(), "avatarId", payload.AvatarId, "avatarType", types.AvatarTypeUser)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	// Verify OTP
	if err := service.otpService.VerifyOTP(payload.Email, types.OtpTypeRegister, payload.OTP); err != nil {
		return responses.ErrorResponse(common.StatusValidationError, "The OTP you entered is incorrect or has expired.", errors.New("invalid otp"))
	}

	// Hash password
	password, err := Validate.HashPassword(payload.Password)
	if err != nil {
		service.container.Logger.Error("auth.service.register-HashPassword", "error", err.Error(), "email", payload.Email, "password", payload.Password)
		return responses.ErrorResponse(common.StatusServerError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	// Create member
	memberId, err := service.memberRepo.Create(rctx, &models.Member{
		Name:     payload.Name,
		Username: payload.Username,
		Email:    payload.Email,
		Password: password,
		AvatarId: payload.AvatarId,
	})
	if err != nil {
		service.container.Logger.Error("auth.service.register-Create", "error", err.Error(), "payload", payload)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	// Create Token
	var accessToken, refreshToken string
	if accessToken, refreshToken, err = service.jwtUtil.GenerateToken(rctx, memberId, types.UserTypeMember, deviceInfo, ip); err != nil {
		service.container.Logger.Error("auth.service.register-GenerateToken", "error", err.Error(), "userId", memberId, "type", types.UserTypeMember)
		return responses.ErrorResponse(common.StatusServerError, "Something went wrong while generating your tokens. Please try again in a moment.", err)
	}

	// Response
	service.container.Logger.Info("auth.service.register.success", "message", "Welcome aboard! 👋 Your account has been created.", "userId", memberId, "type", types.UserTypeMember, "ip", ip)
	return responses.SuccessResponse("Welcome aboard! 👋 Your account has been created.", responses.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func (service *Auth) SendOTP(rctx *requests.RequestContext, payload requests.SentOTP) responses.ServiceResponse {
	var member models.Member

	// Check email
	if err := service.memberRepo.GetMemberByUsernameOrEmail(rctx, "", payload.Email, &member); err != nil {
		if err == gorm.ErrRecordNotFound {
			if payload.Type != types.OtpTypeRegister {
				return responses.ErrorResponse(common.StatusNotFound, "This email address does not exist in our records.", errors.New("email is not exists"))
			}
		} else {
			service.container.Logger.Error("auth.service.sendOTP-GetMemberByUsernameOrEmail", "error", err.Error(), "email", payload.Email)
			return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
		}
	}

	// OTP generate and save
	otp := service.otpService.GenerateOTP()
	if err := service.otpService.SaveOTP(payload.Email, payload.Type, otp); err != nil {
		service.container.Logger.Error("auth.service.sendOTP-SaveOTP", "error", err.Error(), "email", payload.Email, "type", payload.Type)
		return responses.ErrorResponse(common.StatusServerError, "Something went wrong while generating the OTP. Please try again shortly.", err)
	}

	// Send OTP to email
	data := map[string]string{
		"OTP":      otp,
		"UserName": member.Name,
		"Type":     payload.Type.String(),
		"Email":    payload.Email,
	}
	if err := service.emailService.SendEmail(payload.Email, templates.OTP_VERIFICATION_TITLE, templates.OTP_VERIFICATION_TITLE_TEMPLATE, data, []string{}); err != nil {
		service.container.Logger.Error("auth.service.sendOTP-SendEmail", "error", err.Error(), "email", payload.Email, "templateName", templates.OTP_VERIFICATION_TITLE, "templatePath", templates.OTP_VERIFICATION_TITLE_TEMPLATE, "data", data)
		return responses.ErrorResponse(common.StatusServerError, "We couldn’t send the OTP. Please check your email address and try again.", err)
	}

	// Response
	service.container.Logger.Info("auth.service.sendOTP.success", "message", "Done! 🎉 The OTP has been sent to your email address.", "email", payload.Email, "type", payload.Type)
	return responses.SuccessResponse("Done! 🎉 The OTP has been sent to your email address.", nil)
}

func (service *Auth) ResetPassword(rctx *requests.RequestContext, payload requests.ResetPassword, deviceInfo string, ip string) responses.ServiceResponse {
	var member models.Member

	// Check member
	if err := service.memberRepo.GetMemberByUsernameOrEmail(rctx, "", payload.Email, &member); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "We couldn't find an account matching that information.", err)
		}
		service.container.Logger.Error("auth.service.resetPassword-GetMemberByUsernameOrEmail", "error", err.Error(), "email", payload.Email)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	// Validate old and password new password
	if err := Validate.VerifyPassword(member.Password, payload.Password); err == nil {
		return responses.ErrorResponse(common.StatusValidationError, "Your new password must be different from the previous one.", errors.New("password is the same as the previous one"))
	}

	// Verify OTP
	if err := service.otpService.VerifyOTP(payload.Email, types.OtpTypeResetPassword, payload.OTP); err != nil {
		return responses.ErrorResponse(common.StatusValidationError, "The OTP you entered is incorrect or has expired. Please try again.", err)
	}

	// Hash password
	password, err := Validate.HashPassword(payload.Password)
	if err != nil {
		service.container.Logger.Error("auth.service.resetPassword-HashPassword", "error", err.Error(), "email", payload.Email, "password", payload.Password)
		return responses.ErrorResponse(common.StatusServerError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	// Update password
	if err := service.memberRepo.UpdatePasswordByEmail(rctx, payload.Email, password); err != nil {
		service.container.Logger.Error("auth.service.resetPassword-UpdatePasswordByEmail", "error", err.Error(), "email", payload.Email, "password", payload.Password)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	// Create Token
	var accessToken, refreshToken string
	if accessToken, refreshToken, err = service.jwtUtil.GenerateToken(rctx, member.ID, types.UserTypeMember, deviceInfo, ip); err != nil {
		service.container.Logger.Error("auth.service.resetPassword-UpdatePasswordByEmail", "error", err.Error(), "memberId", member.ID, "password", "userType", types.UserTypeMember)
		return responses.ErrorResponse(common.StatusServerError, "We couldn’t generate your login tokens. Please try again shortly.", err)
	}

	// Response
	return responses.SuccessResponse("Your password has been updated successfully. 🎉", responses.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func (service *Auth) GetToken(rctx *requests.RequestContext, payload requests.Token, deviceInfo string, ip string) responses.ServiceResponse {
	var member models.Member

	// Verify refreshToken
	tokenClaims, err := service.jwtUtil.VerifyRefreshToken(rctx, payload.RefreshToken)
	if err != nil {
		return responses.ErrorResponse(common.StatusValidationError, err.Error(), err)
	}

	// Check member
	claims, _ := tokenClaims.(*utils.JWTClaim)
	if err = service.memberRepo.GetMember(rctx, claims.UserId, &member); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "We couldn’t find an account with the provided information.", err)
		}

		service.container.Logger.Error("auth.service.getToken-Get", "error", err.Error(), "userId", claims.UserId)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	// Remove session
	service.authRepo.DeleteSession(rctx, claims.SessionID)

	userType := types.UserTypeMember
	if member.Role == types.MemberRoleSupport {
		userType = types.UserTypeSupportMember
	}

	// Create Token
	var accessToken, refreshToken string
	if accessToken, refreshToken, err = service.jwtUtil.GenerateToken(rctx, member.ID, userType, deviceInfo, ip); err != nil {
		service.container.Logger.Error("auth.service.getToken-GenerateToken", "error", err.Error(), "memberId", member.ID, "userType", userType, "deviceInfo", deviceInfo, "ip", ip)
		return responses.ErrorResponse(common.StatusServerError, "We couldn’t generate your login tokens. Please try again shortly..", err)
	}

	// Response
	return responses.SuccessResponse("You're all set! New tokens have been generated successfully. 🔐", responses.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

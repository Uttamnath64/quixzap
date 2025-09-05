package services

import (
	"github.com/Uttamnath64/quixzap/internal/app/common/types"
	"github.com/Uttamnath64/quixzap/internal/app/utils/requests"
	"github.com/Uttamnath64/quixzap/internal/app/utils/responses"
	"github.com/Uttamnath64/quixzap/pkg/validater"
)

var (
	Validate *validater.Validater
)

type EmailService interface {
	SendEmail(to, subject, templateFile string, data map[string]string, attachments []string) error
}

type OTPService interface {
	GenerateOTP() string
	SaveOTP(email string, otpType types.OtpType, otp string) error
	VerifyOTP(email string, otpType types.OtpType, providedOTP string) error
}

type AuthService interface {
	Login(rctx *requests.RequestContext, payload requests.Login, deviceInfo string, ip string) responses.ServiceResponse
	Register(rctx *requests.RequestContext, payload requests.Register, deviceInfo string, ip string) responses.ServiceResponse
	SendOTP(rctx *requests.RequestContext, payload requests.SentOTP) responses.ServiceResponse
	ResetPassword(rctx *requests.RequestContext, payload requests.ResetPassword, deviceInfo string, ip string) responses.ServiceResponse
	GetToken(rctx *requests.RequestContext, payload requests.Token, deviceInfo string, ip string) responses.ServiceResponse
}

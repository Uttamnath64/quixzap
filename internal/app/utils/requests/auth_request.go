package requests

import (
	"errors"
	"strings"

	"github.com/Uttamnath64/quixzap/internal/app/common/types"
)

type Register struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	AvatarId uint   `json:"avatar_id" binding:"required"`
	OTP      string `json:"otp" binding:"required"`
}

func (r Register) IsValid() error {
	if err := Validate.IsValidName(r.Name); err != nil {
		return err
	}
	if err := Validate.IsValidUsername(r.Username); err != nil {
		return err
	}
	if err := Validate.IsValidEmail(r.Email); err != nil {
		return err
	}
	if err := Validate.IsValidPassword(r.Password); err != nil {
		return err
	}
	if !Validate.IsValidID(r.AvatarId) {
		return errors.New("Please select a valid avatar.")
	}
	if err := Validate.IsValidOTP(r.OTP); err != nil {
		return err
	}
	return nil
}

type SentOTP struct {
	Email string        `json:"email" binding:"required"`
	Type  types.OtpType `json:"type" binding:"required"`
}

func (r SentOTP) IsValid() error {
	if err := Validate.IsValidEmail(r.Email); err != nil {
		return err
	}
	if !r.Type.IsValid() {
		return errors.New("Invalid OTP type.")
	}
	return nil
}

// Reset Password payload
type ResetPassword struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	OTP      string `json:"otp" binding:"required"`
}

func (r ResetPassword) IsValid() error {
	if err := Validate.IsValidEmail(r.Email); err != nil {
		return err
	}
	if err := Validate.IsValidPassword(r.Password); err != nil {
		return err
	}
	if err := Validate.IsValidOTP(r.OTP); err != nil {
		return err
	}
	return nil
}

type Login struct {
	UsernameEmail string `json:"username_email" binding:"required,email|alphanum"`
	Password      string `json:"password" binding:"required,min=6"`
}

func (r Login) IsValid() error {
	emailErr := Validate.IsValidEmail(r.UsernameEmail)
	usernameErr := Validate.IsValidUsername(r.UsernameEmail)
	if emailErr != nil && usernameErr != nil {
		return errors.New("Please enter a valid email address or username.")
	}

	if passErr := Validate.IsValidPassword(r.Password); passErr != nil {
		return passErr
	}
	return nil
}

type Token struct {
	RefreshToken string `json:"refresh_token" binding:"required,min=10"`
}

func (r Token) IsValid() error {
	if strings.TrimSpace(r.RefreshToken) == "" {
		return errors.New("Refresh token is required.")
	}
	return nil
}

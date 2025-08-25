package requests

import (
	"errors"
	"strings"
)

// Login payload
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (r LoginRequest) IsValid() error {
	usernameErr := Validate.IsValidUsername(r.Username)
	if usernameErr != nil {
		return usernameErr
	}

	if passErr := Validate.IsValidPassword(r.Password); passErr != nil {
		return passErr
	}
	return nil
}

// Register payload
type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (r RegisterRequest) IsValid() error {
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
	return nil
}

// Token payload
type TokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func (r TokenRequest) IsValid() error {
	if strings.TrimSpace(r.RefreshToken) == "" {
		return errors.New("Refresh token is required.")
	}
	return nil
}

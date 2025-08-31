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

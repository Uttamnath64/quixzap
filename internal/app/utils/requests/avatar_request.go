package requests

import (
	"errors"
	"strings"

	"github.com/Uttamnath64/quixzap/internal/app/common/types"
)

type AvatarRequest struct {
	Name string           `json:"name" binding:"required"`
	Icon string           `json:"icon" binding:"required"`
	Type types.AvatarType `json:"type" binding:"required"`
}

func (r AvatarRequest) IsValid() error {
	if err := Validate.IsValidName(r.Name); err != nil {
		return err
	}
	if strings.TrimSpace(r.Icon) == "" {
		return errors.New("invalid icon")
	}
	if !r.Type.IsValid() {
		return errors.New("invalid avatar type")
	}

	return nil
}

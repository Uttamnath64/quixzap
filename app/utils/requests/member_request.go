package requests

import "errors"

type UpdateMember struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	AvatarId uint   `json:"avatar_id" binding:"required"`
}

func (r UpdateMember) IsValid() error {
	if err := Validate.IsValidName(r.Name); err != nil {
		return err
	}
	if err := Validate.IsValidUsername(r.Username); err != nil {
		return err
	}
	if !Validate.IsValidID(r.AvatarId) {
		return errors.New("invalid avatar id")
	}
	return nil
}

package models

import "github.com/Uttamnath64/quixzap/app/common/types"

type Avatar struct {
	BaseModel
	Name string           `json:"name" gorm:"not null"`
	Icon string           `json:"icon" gorm:"type:varchar(10);charset:utf8mb4"`
	Type types.AvatarType `json:"type" gorm:"not null"`
}

func (m *Avatar) GetName() string {
	return "avatars"
}

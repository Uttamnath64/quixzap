package models

import "github.com/Uttamnath64/quixzap/app/common/types"

type Member struct {
	BaseModel
	AvatarId    uint             `json:"avatar_id"`
	MemberId    *uint            `json:"customer_id"`
	Name        string           `json:"name" gorm:"not null"`
	Username    string           `json:"username" gorm:"uniqueIndex;not null"`
	Email       string           `json:"email" gorm:"uniqueIndex"`
	Password    string           `json:"-" gorm:"not null"`
	Role        types.MemberRole `json:"role" gorm:"not null; default:1"`
	BusinessIds string           `json:"business_ids" gorm:"type:TEXT; default:'[]'"`
	IsActive    bool             `json:"is_active" gorm:"default:true"`
}

func (m *Member) GetName() string {
	return "members"
}

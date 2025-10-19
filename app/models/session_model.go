package models

import "github.com/Uttamnath64/quixzap/app/common/types"

type Session struct {
	BaseModel
	Theme        int
	UserId       uint           `gorm:"not null"`
	UserType     types.UserType `gorm:"type:VARCHAR(50);not null"`
	DeviceInfo   string         `gorm:"type:TEXT"`
	IPAddress    string         `gorm:"type:VARCHAR(45)"`
	RefreshToken string         `gorm:"type:TEXT"`
	ExpiresAt    int64
}

func (m *Session) GetName() string {
	return "sessions"
}

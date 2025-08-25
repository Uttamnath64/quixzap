package models

import "github.com/google/uuid"

type Session struct {
	ID           uuid.UUID `json:"id" gorm:"type:char(36);primaryKey"`
	AdminId      uuid.UUID `gorm:"not null"`
	DeviceInfo   string    `gorm:"type:TEXT"`
	IPAddress    string    `gorm:"type:VARCHAR(45)"`
	RefreshToken string    `gorm:"type:TEXT"`
	ExpiresAt    int64
}

func (m *Session) GetName() string {
	return "sessions"
}

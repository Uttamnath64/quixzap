package models

import (
	"time"

	"github.com/google/uuid"
)

type Chat struct {
	BaseModel
	UUID       uuid.UUID  `json:"uuid" gorm:"type:char(36);uniqueIndex"`
	BlockedAt  *time.Time `json:"blocked_at"`
	BlockedBy  *uint      `json:"blocked_by"`
	DeviceInfo string     `json:"device_info" gorm:"type:TEXT"`
	IpAddress  string     `json:"ip_address" gorm:"type:VARCHAR(45)"`
}

func (m *Chat) GetName() string {
	return "chats"
}

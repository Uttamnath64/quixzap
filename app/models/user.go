package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:char(36);primaryKey"`
	UserAgent *string   `json:"user_agent"`
	IpAddress *string   `json:"ip_address"`
	IsActive  bool      `json:"is_active" gorm:"default:active"`
	StartedAt time.Time `json:"started_at"`
	EndedAt   time.Time `json:"ended_at"`
}

func (m *User) GetName() string {
	return "users"
}

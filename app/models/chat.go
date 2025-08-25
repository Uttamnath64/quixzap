package models

import (
	"time"

	"github.com/google/uuid"
)

type Chat struct {
	ID        uuid.UUID `json:"id" gorm:"type:char(36);primaryKey"`
	SessionID uuid.UUID `json:"session_id" gorm:"type:char(36);index"`
	Sender    string    `json:"sender" gorm:"type:enum('user','admin')"`
	Content   string    `json:"content" gorm:"type:text"`
	IsRead    bool      `json:"is_read" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at"`
}

func (m *Chat) GetName() string {
	return "chats"
}

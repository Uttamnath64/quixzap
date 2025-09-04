package models

import (
	"time"

	"github.com/google/uuid"
)

type ChatSession struct {
	BaseModel
	BusinessID uint       `json:"business_id"`
	MemberID   uint       `json:"member_id"`
	SessionID  uuid.UUID  `json:"session_id"`
	UserIP     string     `json:"user_ip"`
	UserParams string     `json:"user_params"`
	Status     string     `json:"status"`
	Note       string     `json:"note" gorm:"type:TEXT"`
	ClosedAt   *time.Time `json:"closed_at"`
}

func (m *ChatSession) GetName() string {
	return "chat_sessions"
}

package models

import "github.com/google/uuid"

type ChatMessage struct {
	BaseModel
	ChatID        uint      `json:"chat_id" gorm:"not null"`
	SenderID      uint      `json:"sender_id" gorm:"not null"`
	ChatSessionID uuid.UUID `json:"chat_session_id" gorm:"type:VARCHAR(100); not null"`
	SenderType    string    `json:"sender_type" gorm:"type:VARCHAR(50);not null"`
	Message       string    `json:"message" gorm:"type:TEXT; not null"`
	Status        string    `json:"status" gorm:"type:VARCHAR(50); default:'sent'"`
}

func (m *ChatMessage) GetName() string {
	return "chat_messages"
}

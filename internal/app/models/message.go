package models

import (
	"github.com/google/uuid"
)

type Message struct {
	BaseModel
	ChatUUID uuid.UUID `json:"ChatUUID"`
	ChatId   uint      `json:"chat_id"`
	AdminId  uint      `json:"admin_id"`
	Message  string    `json:"message" gorm:"type:text"`
}

func (m *Message) GetName() string {
	return "messages"
}

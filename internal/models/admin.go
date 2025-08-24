package models

import (
	"time"

	"github.com/google/uuid"
)

type Admin struct {
	ID        uuid.UUID `json:"id" gorm:"type:char(36);primaryKey"`
	Email     string    `json:"email" gorm:"type:char(50);index"`
	Password  string    `json:"password" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at"`
}

func (m *Admin) GetName() string {
	return "chats"
}

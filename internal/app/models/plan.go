package models

type Plan struct {
	BaseModel
	Name         string  `json:"name" gorm:"uniqueIndex;not null"`
	Desscription string  `json:"description" gorm:"type:TEXT"`
	MaxChats     int     `json:"max_chats" gorm:"not null"`
	MaxAdmins    int     `json:"max_admins" gorm:"not null"`
	MaxDomains   int     `json:"max_domains" gorm:"not null"`
	Price        float64 `json:"price" gorm:"not null"`
}

func (m *Plan) GetName() string {
	return "plans"
}

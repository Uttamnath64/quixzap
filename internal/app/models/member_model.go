package models

type Member struct {
	BaseModel
	MemberId    *uint  `json:"customer_id"`
	Name        string `json:"name" gorm:"not null"`
	Username    string `json:"username" gorm:"uniqueIndex;not null"`
	Email       string `json:"email" gorm:"uniqueIndex"`
	Password    string `json:"-" gorm:"not null"`
	Role        string `json:"role" gorm:"type:VARCHAR(50);not null"`
	BusinessIds string `json:"business_ids" gorm:"type:TEXT; default:'[]'"`
	IsActive    bool   `json:"is_active" gorm:"default:true"`
}

func (m *Member) GetName() string {
	return "members"
}

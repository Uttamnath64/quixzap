package models

type Business struct {
	BaseModel
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description" gorm:"type:TEXT"`
	Domain      string `json:"domain" gorm:"uniqueIndex;not null"`
	MemberID    uint   `json:"member_id" gorm:"not null"`
	IsActive    bool   `json:"is_active" gorm:"default:true"`
}

func (m *Business) GetName() string {
	return "businesses"
}

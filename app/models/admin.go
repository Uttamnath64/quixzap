package models

type Admin struct {
	BaseModel
	Name     string `json:"name" gorm:"type:varchar(30);not null"`
	Username string `json:"username" gorm:"type:varchar(20);unique;not null;index"`
	Email    string `json:"email" gorm:"type:varchar(100);unique;not null"`
	Password string `json:"password" gorm:"type:varchar(100);not null"`
	IsActive bool   `json:"is_active" gorm:"default:false"`
}

func (m *Admin) GetName() string {
	return "admins"
}

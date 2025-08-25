package models

type Admin struct {
	BaseModel
	Username string `json:"username" gorm:"type:char(20);index"`
	Password string `json:"password" gorm:"type:text"`
}

func (m *Admin) GetName() string {
	return "chats"
}

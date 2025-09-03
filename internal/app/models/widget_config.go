package models

type WidgetConfig struct {
	BaseModel
	MemberID       uint   `json:"member_id" gorm:"not null"`
	BusinessID     uint   `json:"business_id" gorm:"not null"`
	WidgetName     string `json:"widget_name" gorm:"not null"`
	WelcomeMessage string `json:"welcome_message" gorm:"type:TEXT"`
	Color          string `json:"color"`
	Position       string `json:"position"`
	Width          int    `json:"width"`
	Height         int    `json:"height"`
	RequiredParams string `json:"required_params"`
	IsActive       bool   `json:"is_active" gorm:"default:true"`
}

func (m *WidgetConfig) GetName() string {
	return "widget_configs"
}

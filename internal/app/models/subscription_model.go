package models

import "time"

type Subscription struct {
	BaseModel
	CustomerID  int       `json:"customer_id" gorm:"not null"`
	PlanID      string    `json:"plan_id" gorm:"not null"`
	Status      string    `json:"status" gorm:"type:VARCHAR(50);not null"`
	RenewalDate time.Time `json:"renewal_date" gorm:"not null"`
	Amount      float64   `json:"amount" gorm:"not null"`
	Currency    string    `json:"currency"`
}

func (m *Subscription) GetName() string {
	return "subscriptions"
}

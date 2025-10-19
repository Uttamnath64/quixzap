package models

type BlockedIP struct {
	BaseModel
	CustomerID int    `json:"customer_id"`
	IPAddress  string `json:"ip_address"`
	Reason     string `json:"reason"`
}

func (m *BlockedIP) GetName() string {
	return "blocked_ips"
}

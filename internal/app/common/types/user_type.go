package types

// User, Admin
type UserType int8

const (
	UserTypeUser UserType = iota + 1
	UserTypeCustomer
	UserTypeCustomerSupportPanel
	UserTypeAdmin
)

func (t UserType) String() string {
	names := [...]string{"", "User", "Customer", "Customer Support Panel", "Admin"}
	if !t.IsValid() {
		return "Unknown"
	}
	return names[t]
}

func (t UserType) IsValid() bool {
	return t >= UserTypeUser && t <= UserTypeAdmin
}

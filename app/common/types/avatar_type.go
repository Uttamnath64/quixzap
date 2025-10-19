package types

// Default, Business, User, Status, Panel
type AvatarType int8

const (
	AvatarTypeDefault AvatarType = iota + 1
	AvatarTypeBusiness
	AvatarTypeUser
	AvatarTypeStatus
	AvatarTypePanel
)

func (t AvatarType) String() string {
	names := [...]string{"", "Default", "Business", "User", "Status", "Panel"}
	if !t.IsValid() {
		return "Unknown"
	}
	return names[t]
}

func (t AvatarType) IsValid() bool {
	return t >= AvatarTypeDefault && t <= AvatarTypePanel
}

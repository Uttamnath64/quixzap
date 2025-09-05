package types

// Admin, Support
type MemberRole int8

const (
	MemberRoleAdmin MemberRole = iota + 1
	MemberRoleSupport
)

func (t MemberRole) String() string {
	names := [...]string{"", "Admin", "Support"}
	if !t.IsValid() {
		return "Unknown"
	}
	return names[t]
}

func (t MemberRole) IsValid() bool {
	return t >= MemberRoleAdmin && t <= MemberRoleSupport
}

package types

// "Member", "Support Member", "Admin"
type UserType int8

const (
	UserTypeMember UserType = iota + 1
	UserTypeSupportMember
	UserTypeAdmin
)

func (t UserType) String() string {
	names := [...]string{"", "Member", "Support Member", "Admin"}
	if !t.IsValid() {
		return "Unknown"
	}
	return names[t]
}

func (t UserType) IsValid() bool {
	return t >= UserTypeMember && t <= UserTypeAdmin
}

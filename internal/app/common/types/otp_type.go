package types

// Register, Reset Password
type OtpType int8

const (
	OtpTypeRegister OtpType = iota + 1
	OtpTypeResetPassword
)

func (t OtpType) String() string {
	names := [...]string{"", "Register", "Reset Password"}
	if !t.IsValid() {
		return "Unknown"
	}
	return names[t]
}

func (t OtpType) IsValid() bool {
	return t >= OtpTypeRegister && t <= OtpTypeResetPassword
}

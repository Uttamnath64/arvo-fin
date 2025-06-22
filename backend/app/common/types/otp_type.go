package commonType

// Register, ResetPassword
type OtpType int8

const (
	OtpTypeRegister OtpType = iota + 1
	OtpTypeResetPassword
)

func (t OtpType) String() string {
	return [...]string{"Register", "ResetPassword"}[t]
}

func (t OtpType) IsValid() bool {
	return t >= OtpTypeRegister && t <= OtpTypeResetPassword
}

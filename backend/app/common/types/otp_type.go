package commonType

type OtpType int

const (
	Register OtpType = iota + 1
	ResetPassword
)

func (t OtpType) String() string {
	return [...]string{"Register", "ResetPassword"}[t]
}

func (t OtpType) IsValid() bool {
	return t >= Register && t <= ResetPassword
}

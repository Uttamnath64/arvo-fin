package commonType

// Default, User, Category, Portfolio, RegularPayment
type AvatarType int8

const (
	AvatarTypeDefault AvatarType = iota + 1
	AvatarTypeUser
	AvatarTypeCategory
	AvatarTypePortfolio
	AvatarTypeRegularPayment
)

func (t AvatarType) String() string {
	names := [...]string{"", "Default", "User", "Category", "Portfolio", "RegularPayment"}
	if !t.IsValid() {
		return "Unknown"
	}
	return names[t]
}

func (t AvatarType) IsValid() bool {
	return t >= AvatarTypeDefault && t <= AvatarTypeRegularPayment
}

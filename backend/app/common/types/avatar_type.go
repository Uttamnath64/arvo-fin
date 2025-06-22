package commonType

// Default, User, Category, Portfolio
type AvatarType int8

const (
	AvatarTypeDefault AvatarType = iota + 1
	AvatarTypeUser
	AvatarTypeCategory
	AvatarTypePortfolio
)

func (t AvatarType) String() string {
	names := [...]string{"Default", "User", "Category", "Portfolio"}
	if !t.IsValid() {
		return "Unknown"
	}
	return names[t]
}

func (t AvatarType) IsValid() bool {
	return t >= AvatarTypeDefault && t <= AvatarTypePortfolio
}

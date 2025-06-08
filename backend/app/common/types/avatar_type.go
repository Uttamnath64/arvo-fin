package commonType

// Default, User, Category, Portfolio
type AvatarType int

const (
	AvatarTypeDefault AvatarType = iota + 1
	AvatarTypeUser
	AvatarTypeCategory
	AvatarTypePortfolio
)

func (t AvatarType) String() string {
	return [...]string{"Default", "User", "Category", "Portfolio"}[t]
}

func (t AvatarType) IsValid() bool {
	return t >= AvatarTypeDefault && t <= AvatarTypePortfolio
}

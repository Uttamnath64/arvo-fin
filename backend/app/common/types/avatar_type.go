package commonType

type AvatarType int

const (
	DefaultAvatar AvatarType = iota + 1
	UserAvatar
	CategoryAvatar
	PortfolioAvatar
)

func (t AvatarType) String() string {
	return [...]string{"DefaultAvatar", "UserAvatar", "CategoryAvatar", "PortfolioAvatar"}[t]
}

func (t AvatarType) IsValid() bool {
	return t >= DefaultAvatar && t <= PortfolioAvatar
}

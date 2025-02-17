package commonType

type AvatarType int

const (
	DefaultAvatar AvatarType = iota + 1
)

func (t AvatarType) String() string {
	return [...]string{"Default"}[t]
}

func (t AvatarType) IsValid() bool {
	return t >= DefaultAvatar && t <= DefaultAvatar
}

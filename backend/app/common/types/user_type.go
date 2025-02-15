package commonType

type UserType int

const (
	User UserType = iota + 1
	Admin
)

func (t UserType) String() string {
	return [...]string{"User", "Admin"}[t]
}

func (t UserType) IsValid() bool {
	return t >= User && t <= Admin
}

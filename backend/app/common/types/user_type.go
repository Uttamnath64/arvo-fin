package commonType

// User, Admin
type UserType int

const (
	UserTypeUser UserType = iota + 1
	UserTypeAdmin
)

func (t UserType) String() string {
	return [...]string{"User", "Admin"}[t]
}

func (t UserType) IsValid() bool {
	return t >= UserTypeUser && t <= UserTypeAdmin
}

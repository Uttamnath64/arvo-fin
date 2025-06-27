package commonType

// User, Admin
type UserType int8

const (
	UserTypeUser UserType = iota + 1
	UserTypeAdmin
)

func (t UserType) String() string {
	names := [...]string{"", "User", "Admin"}
	if !t.IsValid() {
		return "Unknown"
	}
	return names[t]
}

func (t UserType) IsValid() bool {
	return t >= UserTypeUser && t <= UserTypeAdmin
}

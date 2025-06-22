package models

type Admin struct {
	BaseModel
	AvatarId uint
	Name     string     `gorm:"type:varchar(30);not null"`
	Email    string     `gorm:"type:varchar(100);unique;not null"`
	Username string     `gorm:"type:varchar(20);unique;not null"`
	Password string     `gorm:"type:varchar(100);not null"`
	Category []Category `gorm:"-"`
}

func (m *Admin) GetName() string {
	return "admins"
}

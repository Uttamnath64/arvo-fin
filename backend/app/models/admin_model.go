package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	AvatarId uint
	Name     string     `gorm:"type:varchar(30);not null"`
	Email    string     `gorm:"type:varchar(100);unique;not null"`
	Username string     `gorm:"type:varchar(20);unique;not null"`
	Password string     `gorm:"type:varchar(100);not null"`
	Avatar   []Avatar   `gorm:"foreignKey:AdminId;constraint:OnDelete:CASCADE;"`
	Category []Category `gorm:"foreignKey:AdminId;constraint:OnDelete:CASCADE;"`
}

func (m *Admin) GetName() string {
	return "admins"
}

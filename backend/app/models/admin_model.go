package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Name     string     `gorm:"column:name;type:varchar(30);not null"`
	Email    string     `gorm:"column:email;type:varchar(100);unique;not null"`
	Username string     `gorm:"column:username;type:varchar(20);unique;not null"`
	Password string     `gorm:"column:password;type:varchar(100);not null"`
	AvatarId uint       `gorm:"column:avatar_id"`
	Avatar   []Avatar   `gorm:"foreignKey:AdminId;constraint:OnDelete:CASCADE;"`
	Category []Category `gorm:"foreignKey:AdminId;constraint:OnDelete:CASCADE;"`
}

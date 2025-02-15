package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string      `gorm:"column:name;type:varchar(30);not null"`
	Email        string      `gorm:"column:email;type:varchar(100);unique;not null"`
	Username     string      `gorm:"column:username;type:varchar(20);unique;not null"`
	MobileNumber string      `gorm:"column:mobile_number;type:varchar(12);unique;not null"`
	Password     string      `gorm:"column:password;type:varchar(100);not null"`
	Portfolio    []Portfolio `gorm:"foreignKey:user_id"`
	Log          []Log       `gorm:"foreignKey:user_id"`
}

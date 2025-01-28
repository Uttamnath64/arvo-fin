package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string      `gorm:"column:name;type:varchar(30);not null"`
	Email        string      `gorm:"column:email;type:varchar(100);unique;not null"`
	MobileNumber string      `gorm:"column:mobileNumber;type:varchar(12);unique;not null"`
	Password     string      `gorm:"column:password;type:varchar(50);not null"`
	Portfolio    []Portfolio `gorm:"foreignKey:userId"`
	Log          []Log       `gorm:"foreignKey:userId"`
}

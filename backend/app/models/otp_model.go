package models

import "gorm.io/gorm"

type OTP struct {
	gorm.Model
	Username string `gorm:"column:username;type:varchar(20);unique;not null"`
	UserType byte   `gorm:"column:user_type"`
	OTP      string `gorm:"column:otp;type:varchar(6);not null"`
}

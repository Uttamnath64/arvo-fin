package models

import "gorm.io/gorm"

type Token struct {
	gorm.Model
	ReferenceId uint   `gorm:"column:referenceId"`
	UserType    byte   `gorm:"column:userType"`
	IP          string `gorm:"column:ip;type:VARCHAR(20)"`
	Token       string `gorm:"column:token;type:VARCHAR(200)"`
	TokenType   string `gorm:"column:name;type:enum('access', 'refresh');not null"`
	ExpiresAt   int64  `gorm:"column:expiresAt"`
}

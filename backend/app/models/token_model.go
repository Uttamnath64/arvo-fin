package models

import "gorm.io/gorm"

type Token struct {
	gorm.Model
	ReferenceId uint   `gorm:"column:reference_id"`
	UserType    byte   `gorm:"column:user_type"`
	IP          string `gorm:"column:ip;type:VARCHAR(20)"`
	Token       string `gorm:"column:token;type:VARCHAR(200)"`
	TokenType   string `gorm:"column:name;type:enum('access', 'refresh');not null"`
	ExpiresAt   int64  `gorm:"column:expires_at"`
}

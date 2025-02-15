package models

import (
	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	ReferenceId uint                `gorm:"column:reference_id"`
	UserType    commonType.UserType `gorm:"column:user_type"`
	IP          string              `gorm:"column:ip;type:VARCHAR(20)"`
	Token       string              `gorm:"column:token;type:VARCHAR(1024)"`
	TokenType   string              `gorm:"column:name;type:enum('access', 'refresh');not null"`
	ExpiresAt   int64               `gorm:"column:expires_at"`
}

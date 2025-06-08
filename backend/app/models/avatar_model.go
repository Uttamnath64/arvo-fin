package models

import (
	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"gorm.io/gorm"
)

type Avatar struct {
	gorm.Model
	Name string                `gorm:"not null"`
	Icon string                `gorm:"type:varchar(10);charset:utf8mb4"`
	Type commonType.AvatarType `gorm:"not null"`
}

func (m *Avatar) GetName() string {
	return "avatars"
}

package models

import (
	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"gorm.io/gorm"
)

type Avatar struct {
	gorm.Model
	Name    string                `gorm:"column:name;not null"`
	Url     string                `gorm:"column:url;type:varchar(500)"`
	Type    commonType.AvatarType `gorm:"column:type;not null"`
	AdminId uint                  `gorm:"column:admin_id"`
}

func (m *Avatar) GetName() string {
	return "avatars"
}

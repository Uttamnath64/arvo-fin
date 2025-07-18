package models

import (
	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
)

type Session struct {
	BaseModel
	Theme        int
	UserId       uint                `gorm:"not null"`
	UserType     commonType.UserType `gorm:"type:VARCHAR(50);not null"`
	DeviceInfo   string              `gorm:"type:TEXT"`
	IPAddress    string              `gorm:"type:VARCHAR(45)"`
	RefreshToken string              `gorm:"type:TEXT"`
	ExpiresAt    int64
}

func (m *Session) GetName() string {
	return "sessions"
}

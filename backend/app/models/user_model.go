package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	AvatarId           uint
	Name               string             `gorm:"type:varchar(30);not null"`
	Email              string             `gorm:"type:varchar(100);unique;not null"`
	Username           string             `gorm:"type:varchar(20);unique;not null"`
	Password           string             `gorm:"type:varchar(100);not null"`
	CurrencyCode       string             `gorm:"not null;default:'INR'"`
	DecimalPlaces      int                `gorm:"not null; default:3"`
	NumberFormat       int                `gorm:"not null; default:1"`
	EmailNotifications bool               `gorm:"default:true"`
	Portfolio          []Portfolio        `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;"`
	TransactionAudit   []TransactionAudit `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;"`
	Account            []Account          `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;"`
	Category           []Category         `gorm:"-"`
}

func (m *User) GetName() string {
	return "users"
}

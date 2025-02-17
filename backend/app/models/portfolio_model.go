package models

import "gorm.io/gorm"

type Portfolio struct {
	gorm.Model
	Name                 string                 `gorm:"column:name;not null"`
	UserId               uint                   `gorm:"column:user_id"`
	AvatarId             uint                   `gorm:"column:avatar_id"`
	Account              []Account              `gorm:"foreignKey:PortfolioId;constraint:OnDelete:CASCADE;"`
	Budget               []Budget               `gorm:"foreignKey:PortfolioId;constraint:OnDelete:CASCADE;"`
	Transaction          []Transaction          `gorm:"foreignKey:PortfolioId;constraint:OnDelete:CASCADE;"`
	RecurringTransaction []RecurringTransaction `gorm:"foreignKey:PortfolioId;constraint:OnDelete:CASCADE;"`
}

func (m Portfolio) GetName() string {
	return "portfolios"
}

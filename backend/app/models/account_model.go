package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	UserId               uint                   `gorm:"column:user_id;not null"`
	PortfolioId          uint                   `gorm:"column:portfolio_id;not null"`
	Name                 string                 `gorm:"column:name;type:varchar(30);not null"`
	Type                 string                 `gorm:"column:type;type:enum('bank', 'credit', 'wallet', 'investment');not null"`
	Balance              float64                `gorm:"column:balance;type:decimal(15,2);default:0.00"`
	Transaction          []Transaction          `gorm:"foreignKey:AccountId;constraint:OnDelete:CASCADE;"`
	RecurringTransaction []RecurringTransaction `gorm:"foreignKey:AccountId;constraint:OnDelete:CASCADE;"`
}

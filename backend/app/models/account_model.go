package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	UserId               uint                   `gorm:"column:userId;not null"`
	PortfolioId          uint                   `gorm:"column:portfolioId;not null"`
	Name                 string                 `gorm:"column:name;type:varchar(30);not null"`
	Type                 string                 `gorm:"column:type;type:enum('bank', 'credit', 'wallet', 'investment');not null"`
	Balance              float64                `gorm:"column:balance;type:decimal(15,2);default:0.00"`
	Transaction          []Transaction          `gorm:"foreignKey:accountId"`
	RecurringTransaction []RecurringTransaction `gorm:"foreignKey:accountId"`
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type RecurringTransaction struct {
	gorm.Model
	AccountId   uint       `gorm:"column:accountId;not null"`
	CategoryId  uint       `gorm:"column:categoryId;not null"`
	PortfolioId uint       `gorm:"column:portfolioId;not null"`
	Amount      float64    `gorm:"column:amount;type:decimal(15,2);not null"`
	Type        string     `gorm:"column:type;type:enum('income', 'expense');not null"`
	Description string     `gorm:"column:description;type:varchar(255)"`
	Frequency   string     `gorm:"column:frequency;type:enum('daily', 'weekly', 'monthly', 'yearly');not null"`
	StartDate   time.Time  `gorm:"column:startDate;not null"`
	EndDate     *time.Time `gorm:"column:endDate;"` // Nullable for indefinite recurring transactions
}

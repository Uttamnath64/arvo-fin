package models

import (
	"time"

	"gorm.io/gorm"
)

type RecurringTransaction struct {
	gorm.Model
	AccountId   uint       `gorm:"column:account_id;not null"`
	CategoryId  uint       `gorm:"column:category_id;not null"`
	PortfolioId uint       `gorm:"column:portfolio_id;not null"`
	Amount      float64    `gorm:"column:amount;type:decimal(15,2);not null"`
	Type        string     `gorm:"column:type;type:enum('income', 'expense');not null"`
	Description string     `gorm:"column:description;type:varchar(255)"`
	Frequency   string     `gorm:"column:frequency;type:enum('daily', 'weekly', 'monthly', 'yearly');not null"`
	StartDate   time.Time  `gorm:"column:start_date;not null"`
	EndDate     *time.Time `gorm:"column:end_date;"` // Nullable for indefinite recurring transactions
}

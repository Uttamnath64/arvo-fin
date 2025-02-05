package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	AccountId         uint    `gorm:"column:account_id;not null"`
	TransferAccountId *uint   `gorm:"column:transfer_account_id;"`
	CategoryId        uint    `gorm:"column:category_id;not null"`
	PortfolioId       uint    `gorm:"column:portfolio_id;not null"`
	Amount            float64 `gorm:"column:amount;type:decimal(15,2);not null"`
	Type              string  `gorm:"column:type;type:enum('income', 'expense', 'transfer');not null"`
	Description       string  `gorm:"column:description;type:varchar(255)"`
	Log               []Log   `gorm:"foreignKey:transaction_id"`
}

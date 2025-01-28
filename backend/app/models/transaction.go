package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	AccountId         uint    `gorm:"column:accountId;not null"`
	TransferAccountId *uint   `gorm:"column:transferAccountId;"`
	CategoryId        uint    `gorm:"column:categoryId;not null"`
	PortfolioId       uint    `gorm:"column:portfolioId;not null"`
	Amount            float64 `gorm:"column:amount;type:decimal(15,2);not null"`
	Type              string  `gorm:"column:type;type:enum('income', 'expense', 'transfer');not null"`
	Description       string  `gorm:"column:description;type:varchar(255)"`
	Log               []Log   `gorm:"foreignKey:transactionId"`
}

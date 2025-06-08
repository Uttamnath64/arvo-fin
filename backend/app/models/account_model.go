package models

import (
	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	UserId               uint                   `gorm:"not null"`
	PortfolioId          uint                   `gorm:"not null"`
	Name                 string                 `gorm:"type:varchar(30);not null"`
	Type                 commonType.AccountType `gorm:"not null"`
	Balance              float64                `gorm:"type:decimal(15,2);default:0.00"`
	Transaction          []Transaction          `gorm:"foreignKey:AccountId;constraint:OnDelete:CASCADE;"`
	RecurringTransaction []RecurringTransaction `gorm:"foreignKey:AccountId;constraint:OnDelete:CASCADE;"`
}

func (m *Account) GetName() string {
	return "accounts"
}

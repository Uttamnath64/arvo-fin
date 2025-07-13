package models

import (
	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
)

type Account struct {
	BaseModel
	UserId         uint                   `json:"user_id" gorm:"not null"`
	PortfolioId    uint                   `json:"portfolio_id" gorm:"not null"`
	AvatarId       uint                   `json:"avatar_id" gorm:"not null"`
	Name           string                 `json:"name" gorm:"type:varchar(30);not null"`
	Type           commonType.AccountType `json:"type" gorm:"not null"`
	CurrencyCode   string                 `json:"currency_code" gorm:"not null;default:'INR'"`
	OpeningBalance float64                `json:"opening_balance" gorm:"type:decimal(15,2);default:0.00"`
	Note           string                 `json:"note" gorm:"type:text"`

	// Relationships
	Avatar   *Avatar   `json:"avatar,omitempty" gorm:"foreignKey:AvatarId"`
	Currency *Currency `json:"currency,omitempty" gorm:"foreignKey:CurrencyCode;references:Code"`

	Transaction          []Transaction          `gorm:"foreignKey:AccountId;constraint:OnDelete:CASCADE;"`
	RecurringTransaction []RecurringTransaction `gorm:"foreignKey:AccountId;constraint:OnDelete:CASCADE;"`
}

func (m *Account) GetName() string {
	return "accounts"
}

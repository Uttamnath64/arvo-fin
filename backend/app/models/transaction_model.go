package models

import commonType "github.com/Uttamnath64/arvo-fin/app/common/types"

type Transaction struct {
	BaseModel
	UserId            uint                       `json:"user_id"`
	TransferAccountId *uint                      `json:"transfer_account_id"`
	AccountId         uint                       `json:"account_id" gorm:"not null"`
	CategoryId        uint                       `json:"category_id" gorm:"not null"`
	PortfolioId       uint                       `json:"portfolio_id" gorm:"not null"`
	Amount            float64                    `json:"amount" gorm:"type:decimal(15,2);not null"`
	Type              commonType.TransactionType `json:"type" gorm:"not null"`
	Note              string                     `json:"note" gorm:"type:varchar(255)"`

	// Relationships
	Account         *Account  `json:"account,omitempty" gorm:"foreignKey:AccountId"`
	TransferAccount *Account  `json:"transfer_account,omitempty" gorm:"foreignKey:TransferAccountId"`
	Category        *Category `json:"category,omitempty" gorm:"foreignKey:CategoryId"`

	TransactionAudit []TransactionAudit `gorm:"foreignKey:TransactionId;constraint:OnDelete:CASCADE;"`
}

func (m *Transaction) GetName() string {
	return "transactions"
}

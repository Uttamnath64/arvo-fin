package models

import (
	"time"

	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
)

type RegularPayments struct {
	BaseModel
	UserId           uint                             `json:"user_id" gorm:"not null"`
	AccountId        uint                             `json:"account_id" gorm:"not null"`
	AvatarId         uint                             `json:"avatar_id" gorm:"not null"`
	Name             string                           `json:"name" gorm:"not null"`
	PaymentFrequency *commonType.PaymentFrequency     `json:"payment_frequency"`
	PaymentDate      *time.Time                       `json:"payment_date"`
	CategoryId       uint                             `json:"category_id" gorm:"not null"`
	PortfolioId      uint                             `json:"portfolio_id" gorm:"not null"`
	Status           *commonType.RegularPaymentStatus `json:"status"`
	Direction        commonType.TransactionDirection  `json:"direction"`
	Amount           float64                          `json:"amount" gorm:"type:decimal(15,2);not null"`
	Note             string                           `json:"note" gorm:"type:varchar(255)"`
	IsInstallment    bool                             `json:"is_installment"`

	// Relationships
	Account  *Account  `json:"account,omitempty" gorm:"foreignKey:AccountId"`
	Avatar   *Avatar   `json:"avatar,omitempty" gorm:"foreignKey:AvatarId"`
	Category *Category `json:"category,omitempty" gorm:"foreignKey:CategoryId"`

	TransactionAudit []TransactionAudit `gorm:"foreignKey:TransactionId;constraint:OnDelete:CASCADE;"`
}

func (m *RegularPayments) GetName() string {
	return "regular_payments"
}

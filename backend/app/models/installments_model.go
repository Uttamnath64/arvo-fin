package models

import (
	"time"
)

type Installments struct {
	BaseModel
	UserId           uint       `json:"user_id" gorm:"not null"`
	AccountId        uint       `json:"account_id" gorm:"not null"`
	RegularPaymentId uint       `json:"regular_payment_id" gorm:"not null"`
	DueDate          time.Time  `json:"due_date" gorm:"not null"`
	Amount           float64    `json:"amount" gorm:"type:decimal(15,2);not null"`
	Note             string     `json:"note" gorm:"type:varchar(255)"`
	Paid             bool       `json:"paid"`
	PaidAt           *time.Time `json:"paid_at"`

	// Relationships
	Account *Account `json:"account,omitempty" gorm:"foreignKey:AccountId"`

	TransactionAudit []TransactionAudit `gorm:"foreignKey:TransactionId;constraint:OnDelete:CASCADE;"`
}

func (m *Installments) GetName() string {
	return "installments"
}

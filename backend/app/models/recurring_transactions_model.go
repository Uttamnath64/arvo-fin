package models

import (
	"time"

	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"gorm.io/gorm"
)

type RecurringTransaction struct {
	gorm.Model
	AccountId   uint                       `gorm:"not null"`
	CategoryId  uint                       `gorm:"not null"`
	PortfolioId uint                       `gorm:"not null"`
	Amount      float64                    `gorm:"type:decimal(15,2);not null"`
	Type        commonType.TransactionType `gorm:"not null"`
	Description string                     `gorm:"type:varchar(255)"`
	Frequency   commonType.FrequencyType   `gorm:"not null"`
	StartDate   time.Time                  `gorm:"not null"`
	EndDate     *time.Time
}

func (m *RecurringTransaction) GetName() string {
	return "recurring_transactions"
}

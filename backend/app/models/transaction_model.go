package models

type Transaction struct {
	BaseModel
	TransferAccountId *uint
	AccountId         uint               `gorm:"not null"`
	CategoryId        uint               `gorm:"not null"`
	PortfolioId       uint               `gorm:"not null"`
	Amount            float64            `gorm:"type:decimal(15,2);not null"`
	Type              string             `gorm:"type:enum('income', 'expense', 'transfer');not null"`
	Description       string             `gorm:"type:varchar(255)"`
	TransactionAudit  []TransactionAudit `gorm:"foreignKey:TransactionId;constraint:OnDelete:CASCADE;"`
}

func (m *Transaction) GetName() string {
	return "transactions"
}

package models

type TransactionAudit struct {
	BaseModel
	UserId        *uint
	TransactionId *uint
	Action        string `gorm:"type:varchar(255);not null"`
}

func (m *TransactionAudit) GetName() string {
	return "transaction_audits"
}

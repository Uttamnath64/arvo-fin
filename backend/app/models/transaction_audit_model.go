package models

import "gorm.io/gorm"

type TransactionAudit struct {
	gorm.Model
	UserId        *uint
	TransactionId *uint
	Action        string `gorm:"type:varchar(255);not null"`
}

func (m *TransactionAudit) GetName() string {
	return "transaction_audits"
}

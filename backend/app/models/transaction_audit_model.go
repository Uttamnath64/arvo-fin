package models

import "gorm.io/gorm"

type TransactionAudit struct {
	gorm.Model
	UserId        *uint  `gorm:"column:user_id;"`        // Nullable for system actions
	TransactionId *uint  `gorm:"column:transaction_id;"` // Nullable for system actions
	Action        string `gorm:"column:action;type:varchar(255);not null"`
}

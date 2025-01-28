package models

import "gorm.io/gorm"

type Log struct {
	gorm.Model
	UserId        *uint  `gorm:"column:userId;"`        // Nullable for system actions
	TransactionId *uint  `gorm:"column:transactionId;"` // Nullable for system actions
	Action        string `gorm:"column:action;type:varchar(255);not null"`
}

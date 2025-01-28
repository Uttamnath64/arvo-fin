package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	UserID      *uint  `gorm:"column:userId;"`
	PortfolioId uint   `gorm:"column:portfolioId;not null"`
	Name        string `gorm:"column:name;type:varchar(100);not null"`
	Type        string `gorm:"column:type;type:enum('income', 'expense');not null"`
}

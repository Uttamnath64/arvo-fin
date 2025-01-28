package models

import "gorm.io/gorm"

type Portfolio struct {
	gorm.Model
	Name                 string                 `gorm:"column:name;not null"`
	UserId               string                 `gorm:"column:userId"`
	Account              []Account              `gorm:"foreignKey:portfolioId"`
	Budget               []Budget               `gorm:"foreignKey:portfolioId"`
	Transaction          []Transaction          `gorm:"foreignKey:portfolioId"`
	RecurringTransaction []RecurringTransaction `gorm:"foreignKey:portfolioId"`
}

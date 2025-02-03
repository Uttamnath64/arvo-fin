package models

import "gorm.io/gorm"

type Portfolio struct {
	gorm.Model
	Name                 string                 `gorm:"column:name;not null"`
	UserId               string                 `gorm:"column:user_id"`
	Account              []Account              `gorm:"foreignKey:portfolio_id"`
	Budget               []Budget               `gorm:"foreignKey:portfolio_id"`
	Transaction          []Transaction          `gorm:"foreignKey:portfolio_id"`
	RecurringTransaction []RecurringTransaction `gorm:"foreignKey:portfolio_id"`
}

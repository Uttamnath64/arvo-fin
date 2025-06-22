package models

import (
	"time"
)

type Budget struct {
	BaseModel
	UserId      uint      `gorm:"not null"`
	CategoryId  uint      `gorm:"not null"`
	PortfolioId uint      `gorm:"not null"`
	Amount      float64   `gorm:"type:decimal(15,2);not null"`
	StartDate   time.Time `gorm:"not null"`
	EndDate     time.Time `gorm:"not null"`
}

func (m *Budget) GetName() string {
	return "budgets"
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type Budget struct {
	gorm.Model
	UserId      uint      `gorm:"column:userId;not null"`
	CategoryId  uint      `gorm:"column:categoryId;not null"`
	PortfolioId uint      `gorm:"column:portfolioId;not null"`
	Amount      float64   `gorm:"column:amount;type:decimal(15,2);not null"`
	StartDate   time.Time `gorm:"column:startDate;not null"`
	EndDate     time.Time `gorm:"column:endDate;not null"`
}

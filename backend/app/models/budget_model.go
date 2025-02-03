package models

import (
	"time"

	"gorm.io/gorm"
)

type Budget struct {
	gorm.Model
	UserId      uint      `gorm:"column:user_id;not null"`
	CategoryId  uint      `gorm:"column:category_id;not null"`
	PortfolioId uint      `gorm:"column:portfolio_id;not null"`
	Amount      float64   `gorm:"column:amount;type:decimal(15,2);not null"`
	StartDate   time.Time `gorm:"column:start_date;not null"`
	EndDate     time.Time `gorm:"column:end_date;not null"`
}

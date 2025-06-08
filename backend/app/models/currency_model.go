package models

type Currency struct {
	Code   string `gorm:"primaryKey;size:10"`
	Name   string `gorm:"size:100;not null"`
	Symbol string `gorm:"size:10"`
}

func (m Currency) GetName() string {
	return "currencies"
}

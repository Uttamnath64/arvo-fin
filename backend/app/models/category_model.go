package models

import (
	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
)

type Category struct {
	BaseModel
	SourceID     uint
	SourceType   commonType.UserType
	PortfolioId  *uint
	CopiedFromID *uint
	Name         string                     `gorm:"type:varchar(100);not null"`
	Type         commonType.TransactionType `gorm:"not null"`
}

func (m *Category) GetName() string {
	return "categories"
}

package models

import (
	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
)

type Category struct {
	BaseModel
	SourceId     uint                       `json:"source_id" gorm:"not null"`
	AvatarId     uint                       `json:"avatar_id" gorm:"not null"`
	SourceType   commonType.UserType        `json:"source_type" gorm:"not null"`
	PortfolioId  *uint                      `json:"portfolio_id"`
	CopiedFromId *uint                      `json:"copied_from_id"`
	Name         string                     `json:"name" gorm:"type:varchar(100);not null"`
	Type         commonType.TransactionType `json:"type" gorm:"not null"`

	// Relationships
	Avatar *Avatar `json:"avatar,omitempty" gorm:"foreignKey:AvatarId"`
}

func (m *Category) GetName() string {
	return "categories"
}

package models

type Portfolio struct {
	BaseModel
	UserId   uint   `json:"user_id"`
	AvatarId uint   `json:"avatar_id"`
	Name     string `json:"name" gorm:"not null"`

	// Relationships
	Avatar *Avatar `json:"avatar,omitempty" gorm:"foreignKey:AvatarId"`

	Account     []Account     `gorm:"foreignKey:PortfolioId;constraint:OnDelete:CASCADE;"`
	Budget      []Budget      `gorm:"foreignKey:PortfolioId;constraint:OnDelete:CASCADE;"`
	Transaction []Transaction `gorm:"foreignKey:PortfolioId;constraint:OnDelete:CASCADE;"`
}

func (m Portfolio) GetName() string {
	return "portfolios"
}

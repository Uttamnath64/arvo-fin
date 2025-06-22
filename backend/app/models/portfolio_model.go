package models

type Portfolio struct {
	BaseModel
	UserId               uint
	AvatarId             uint
	Name                 string                 `gorm:"not null"`
	Account              []Account              `gorm:"foreignKey:PortfolioId;constraint:OnDelete:CASCADE;"`
	Budget               []Budget               `gorm:"foreignKey:PortfolioId;constraint:OnDelete:CASCADE;"`
	Transaction          []Transaction          `gorm:"foreignKey:PortfolioId;constraint:OnDelete:CASCADE;"`
	RecurringTransaction []RecurringTransaction `gorm:"foreignKey:PortfolioId;constraint:OnDelete:CASCADE;"`
}

func (m Portfolio) GetName() string {
	return "portfolios"
}

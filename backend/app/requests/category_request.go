package requests

import (
	"errors"

	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
)

type CategoryRequest struct {
	PortfolioId uint                       `json:"portfolio_id" binding:"required"`
	AvatarId    uint                       `json:"avatar_id" binding:"required"`
	Type        commonType.TransactionType `json:"type" binding:"required"`
	Name        string                     `json:"name" binding:"required"`
}

func (r CategoryRequest) IsValid() error {
	if !Validate.IsValidID(r.PortfolioId) {
		return errors.New("Invalid portfolio id!")
	}
	if !Validate.IsValidID(r.AvatarId) {
		return errors.New("Invalid avater id!")
	}
	if err := Validate.IsValidName(r.Name); err != nil {
		return err
	}
	if !r.Type.IsValid() {
		return errors.New("Invalid transaction type!")
	}
	return nil
}

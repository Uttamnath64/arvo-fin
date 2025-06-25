package requests

import (
	"errors"
)

// Portfolio payload
type PortfolioRequest struct {
	Name     string `json:"name" binding:"required"`
	AvatarId uint   `json:"avatar_id" binding:"required"`
}

func (r PortfolioRequest) IsValid() error {
	if err := Validate.IsValidName(r.Name); err != nil {
		return err
	}
	if !Validate.IsValidID(r.AvatarId) {
		return errors.New("invalid avatar id")
	}
	return nil
}

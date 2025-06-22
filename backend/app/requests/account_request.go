package requests

import (
	"errors"
	"strings"

	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
)

type AccountRequest struct {
	PortfolioId    uint                   `json:"portfolio_id" binding:"required"`
	AvatarId       uint                   `json:"avatar_id" binding:"required"`
	Name           string                 `json:"name" binding:"required"`
	Type           commonType.AccountType `json:"type" binding:"required"`
	CurrencyCode   string                 `json:"currency_code" binding:"required"`
	OpeningBalance float64                `json:"opening_balance,omitempty" binding:"required"`
	Note           string                 `json:"note,omitempty"`
}

func (r AccountRequest) IsValid() error {
	if err := Validate.IsValidName(r.Name); err != nil {
		return err
	}
	if !r.Type.IsValid() {
		return errors.New("invalid account type")
	}
	if strings.TrimSpace(r.CurrencyCode) == "" {
		return errors.New("invalid currency code")
	}
	if !Validate.IsValidID(r.AvatarId) {
		return errors.New("invalid avatar id")
	}
	if !Validate.IsValidID(r.PortfolioId) {
		return errors.New("invalid portfolio id")
	}
	return nil
}

type AccountUpdateRequest struct {
	AvatarId     uint                   `json:"avatar_id" binding:"required"`
	Name         string                 `json:"name" binding:"required"`
	Type         commonType.AccountType `json:"type" binding:"required"`
	CurrencyCode string                 `json:"currency_code" binding:"required"`
	Note         string                 `json:"note,omitempty"`
}

func (r AccountUpdateRequest) IsValid() error {
	if err := Validate.IsValidName(r.Name); err != nil {
		return err
	}
	if !r.Type.IsValid() {
		return errors.New("invalid account type")
	}
	if strings.TrimSpace(r.CurrencyCode) == "" {
		return errors.New("invalid currency code")
	}
	if !Validate.IsValidID(r.AvatarId) {
		return errors.New("invalid avatar id")
	}
	return nil
}

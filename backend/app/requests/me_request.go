package requests

import (
	"errors"
	"strings"

	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
)

type MeRequest struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	AvatarId uint   `json:"avatar_id" binding:"required"`
}

func (r MeRequest) IsValid() error {
	if err := Validate.IsValidName(r.Name); err != nil {
		return err
	}
	if err := Validate.IsValidUsername(r.Username); err != nil {
		return err
	}
	if !Validate.IsValidID(r.AvatarId) {
		return errors.New("Invalid avatar id!")
	}
	return nil
}

type SettingsRequest struct {
	CurrencyCode       string                   `json:"currency_code" binding:"required"`
	DecimalPlaces      commonType.DecimalPlaces `json:"decimal_places" binding:"required"`
	NumberFormat       commonType.NumberFormat  `json:"number_format" binding:"required"`
	RemindEveryday     bool                     `json:"remind_everyday" binding:"required"`
	MonthlyReportEmail bool                     `json:"monthly_report_email" binding:"required"`
}

func (r SettingsRequest) IsValid() error {
	if strings.TrimSpace(r.CurrencyCode) == "" {
		return errors.New("Invalid currency code!")
	}
	if !r.DecimalPlaces.IsValid() {
		return errors.New("Invalid decimal places!")
	}
	if !r.NumberFormat.IsValid() {
		return errors.New("Invalid number format!")
	}
	return nil
}

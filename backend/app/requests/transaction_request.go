package requests

import (
	"errors"
	"time"

	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
)

type TransactionRequest struct {
	TransferAccountId *uint                      `json:"transfer_account_id,omitempty"`
	AccountId         uint                       `json:"account_id" binding:"required"`
	CategoryId        uint                       `json:"category_id" binding:"required"`
	PortfolioId       uint                       `json:"portfolio_id" binding:"required"`
	Amount            float64                    `json:"amount" binding:"required"`
	Type              commonType.TransactionType `json:"type" binding:"required"`
	Note              string                     `json:"note,omitempty"`
}

func (r TransactionRequest) IsValid() error {

	if !Validate.IsValidID(r.AccountId) {
		return errors.New("invalid account id")
	}

	if !Validate.IsValidID(r.CategoryId) {
		return errors.New("invalid category id")
	}

	if !Validate.IsValidID(r.PortfolioId) {
		return errors.New("invalid portfolio id")
	}
	if r.Amount <= 0 {
		return errors.New("invalid amount")
	}
	if !r.Type.IsValid() {
		return errors.New("invalid account type")
	}
	return nil
}

type TransactionQuery struct {
	UserId      uint
	PortfolioId uint
	AccountId   uint
	CategoryId  uint
	DateFrom    time.Time
	DateTo      time.Time
	Search      string
	Type        *commonType.TransactionType
	Order       commonType.OrderType
}

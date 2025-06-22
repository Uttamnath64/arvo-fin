package services

import (
	"errors"

	"github.com/Uttamnath64/arvo-fin/app/common"
	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/repository"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/responses"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"gorm.io/gorm"
)

type Account struct {
	container     *storage.Container
	repoAccount   repository.AccountRepository
	repoAvatar    repository.AvatarRepository
	repoPortfolio repository.PortfolioRepository
	repoCurrency  repository.CurrencyRepository
}

func NewAccount(container *storage.Container) *Account {
	return &Account{
		container:     container,
		repoAccount:   repository.NewAccount(container),
		repoAvatar:    repository.NewAvatar(container),
		repoPortfolio: repository.NewPortfolio(container),
		repoCurrency:  repository.NewCurrency(container),
	}
}

func (service *Account) Get(id uint) responses.ServiceResponse {

	response, err := service.repoAccount.Get(id)
	if err == gorm.ErrRecordNotFound {
		return responses.ServiceResponse{
			StatusCode: common.StatusNotFound,
			Message:    "Account not found!",
			Error:      err,
		}
	}

	if err != nil {
		service.container.Logger.Error("account.service.get", err.Error(), id)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      err,
		}
	}

	// Response
	return responses.ServiceResponse{
		StatusCode: common.StatusSuccess,
		Message:    "Account records found!",
		Data:       response,
	}
}

func (service *Account) GetList(portfolioId, userId uint) responses.ServiceResponse {
	response, err := service.repoAccount.GetList(portfolioId, userId)
	if err == gorm.ErrRecordNotFound {
		return responses.ServiceResponse{
			StatusCode: common.StatusNotFound,
			Message:    "Accounts not found!",
			Error:      err,
		}
	}
	if err != nil {
		service.container.Logger.Error("account.service.get-list", err.Error(), userId)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      err,
		}
	}

	// Response
	return responses.ServiceResponse{
		StatusCode: common.StatusSuccess,
		Message:    "Accounts found!",
		Data:       response,
	}
}

func (service *Account) AccountTypes() responses.ServiceResponse {
	type AccountTypeSeed struct {
		Name string
		Icon *models.Avatar
	}

	accountTypes := map[commonType.AccountType]AccountTypeSeed{
		commonType.AccountTypeBank: {
			Name: commonType.AccountTypeBank.String(),
			Icon: service.repoAvatar.GetByNameAndType("Bank", commonType.AvatarTypePortfolio),
		},
		commonType.AccountTypeCash: {
			Name: commonType.AccountTypeCash.String(),
			Icon: service.repoAvatar.GetByNameAndType("Cash", commonType.AvatarTypePortfolio),
		},
		commonType.AccountTypeWallet: {
			Name: commonType.AccountTypeWallet.String(),
			Icon: service.repoAvatar.GetByNameAndType("Wallet", commonType.AvatarTypePortfolio),
		},
		commonType.AccountTypeCredit: {
			Name: commonType.AccountTypeCredit.String(),
			Icon: service.repoAvatar.GetByNameAndType("Credit", commonType.AvatarTypePortfolio),
		},
		commonType.AccountTypeLoan: {
			Name: commonType.AccountTypeLoan.String(),
			Icon: service.repoAvatar.GetByNameAndType("Loan", commonType.AvatarTypePortfolio),
		},
		commonType.AccountTypeInvestment: {
			Name: commonType.AccountTypeInvestment.String(),
			Icon: service.repoAvatar.GetByNameAndType("Chart", commonType.AvatarTypePortfolio),
		},
	}

	// Response
	return responses.ServiceResponse{
		StatusCode: common.StatusSuccess,
		Message:    "Account type found!",
		Data:       accountTypes,
	}
}

func (service *Account) Create(userId uint, payload requests.AccountRequest) responses.ServiceResponse {
	ok, err := service.repoPortfolio.UserPortfolioExists(payload.PortfolioId, userId)
	if err != nil {
		service.container.Logger.Error("account.service.create", err.Error(), payload)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later!",
			Error:      err,
		}
	}
	if !ok {
		return responses.ServiceResponse{
			StatusCode: common.StatusNotFound,
			Message:    "Portfolio not found with this user!",
			Error:      errors.New("portfolio not found with this user"),
		}
	}

	ok, err = service.repoAvatar.AvatarByTypeExists(payload.AvatarId, commonType.AvatarTypePortfolio)
	if err != nil {
		service.container.Logger.Error("account.service.create", err.Error(), payload)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later!",
			Error:      err,
		}
	}
	if !ok {
		return responses.ServiceResponse{
			StatusCode: common.StatusNotFound,
			Message:    "Avatar not found!",
			Error:      errors.New("avatar not found"),
		}
	}

	ok, err = service.repoCurrency.CodeExists(payload.CurrencyCode)
	if err != nil {
		service.container.Logger.Error("account.service.create", err.Error(), payload)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later!",
			Error:      err,
		}
	}
	if !ok {
		return responses.ServiceResponse{
			StatusCode: common.StatusNotFound,
			Message:    "Currency not found!",
			Error:      err,
		}
	}

	id, err := service.repoAccount.Create(models.Account{
		UserId:         userId,
		PortfolioId:    payload.PortfolioId,
		AvatarId:       payload.AvatarId,
		Name:           payload.Name,
		Type:           payload.Type,
		CurrencyCode:   payload.CurrencyCode,
		OpeningBalance: payload.OpeningBalance,
		Note:           payload.Note,
	})

	if err != nil {
		service.container.Logger.Error("account.service.create", err.Error(), payload)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later!",
			Error:      err,
		}
	}

	response, _ := service.repoAccount.Get(id)
	return responses.ServiceResponse{
		StatusCode: common.StatusSuccess,
		Message:    "Account is created!",
		Data:       response,
	}
}

func (service *Account) Update(id, userId uint, payload requests.AccountUpdateRequest) responses.ServiceResponse {

	ok, err := service.repoAvatar.AvatarByTypeExists(payload.AvatarId, commonType.AvatarTypePortfolio)
	if err != nil {
		service.container.Logger.Error("account.service.update", err.Error(), payload)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later!",
			Error:      err,
		}
	}
	if !ok {
		return responses.ServiceResponse{
			StatusCode: common.StatusNotFound,
			Message:    "Avatar not found!",
			Error:      errors.New("avatar not found"),
		}
	}

	ok, err = service.repoCurrency.CodeExists(payload.CurrencyCode)
	if err != nil {
		service.container.Logger.Error("account.service.update", err.Error(), payload, userId)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later!",
			Error:      err,
		}
	}
	if !ok {
		return responses.ServiceResponse{
			StatusCode: common.StatusNotFound,
			Message:    "Currency not found!",
			Error:      err,
		}
	}

	ok, err = service.repoAccount.Update(id, userId, payload)
	if err != nil {
		service.container.Logger.Error("account.service.update", err.Error(), id, userId, payload)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    err.Error(),
			Error:      err,
		}
	}
	if !ok {
		return responses.ServiceResponse{
			StatusCode: common.StatusNotFound,
			Message:    "Account not updated!",
			Error:      errors.New("account not updated"),
		}
	}

	response, _ := service.repoAccount.Get(id)
	return responses.ServiceResponse{
		StatusCode: common.StatusSuccess,
		Message:    "Account record updated!",
		Data:       response,
	}
}

func (service *Account) Delete(id, userId uint) responses.ServiceResponse {
	ok, err := service.repoAccount.Delete(id, userId)
	if err != nil {
		service.container.Logger.Error("account.service.deleted", err.Error(), id, userId)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    err.Error(),
			Error:      err,
		}
	}
	if !ok {
		return responses.ServiceResponse{
			StatusCode: common.StatusNotFound,
			Message:    "Account not deleted!",
			Error:      errors.New("account not deleted"),
		}
	}

	return responses.ServiceResponse{
		StatusCode: common.StatusSuccess,
		Message:    "Account deleted!",
	}
}

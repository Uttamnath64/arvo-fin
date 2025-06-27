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

func (service *Account) Get(rctx *requests.RequestContext, id uint) responses.ServiceResponse {

	response, err := service.repoAccount.Get(rctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Account not found!", err)
		}

		service.container.Logger.Error("account.appService.get-Get", "error", err.Error(), "id", id)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong. Please try again later.", err)
	}

	// Response
	return responses.SuccessResponse("Account records found!", response)
}

func (service *Account) GetList(rctx *requests.RequestContext, portfolioId, userId uint) responses.ServiceResponse {
	response, err := service.repoAccount.GetList(rctx, portfolioId, userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Accounts not found!", err)
		}

		service.container.Logger.Error("account.appService.getList-GetList", "error", err.Error(), "portfolioId", portfolioId, "userId", userId)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong. Please try again later.", err)
	}

	// Response
	return responses.SuccessResponse("Accounts found!", response)
}

func (service *Account) AccountTypes(rctx *requests.RequestContext) responses.ServiceResponse {
	type AccountTypeSeed struct {
		Name string
		Icon *models.Avatar
	}

	accountTypes := map[commonType.AccountType]AccountTypeSeed{
		commonType.AccountTypeBank: {
			Name: commonType.AccountTypeBank.String(),
			Icon: service.repoAvatar.GetByNameAndType(rctx, "Bank", commonType.AvatarTypePortfolio),
		},
		commonType.AccountTypeCash: {
			Name: commonType.AccountTypeCash.String(),
			Icon: service.repoAvatar.GetByNameAndType(rctx, "Cash", commonType.AvatarTypePortfolio),
		},
		commonType.AccountTypeWallet: {
			Name: commonType.AccountTypeWallet.String(),
			Icon: service.repoAvatar.GetByNameAndType(rctx, "Wallet", commonType.AvatarTypePortfolio),
		},
		commonType.AccountTypeCredit: {
			Name: commonType.AccountTypeCredit.String(),
			Icon: service.repoAvatar.GetByNameAndType(rctx, "Credit", commonType.AvatarTypePortfolio),
		},
		commonType.AccountTypeLoan: {
			Name: commonType.AccountTypeLoan.String(),
			Icon: service.repoAvatar.GetByNameAndType(rctx, "Loan", commonType.AvatarTypePortfolio),
		},
		commonType.AccountTypeInvestment: {
			Name: commonType.AccountTypeInvestment.String(),
			Icon: service.repoAvatar.GetByNameAndType(rctx, "Chart", commonType.AvatarTypePortfolio),
		},
	}

	// Response
	return responses.SuccessResponse("Account type found!", accountTypes)
}

func (service *Account) Create(rctx *requests.RequestContext, userId uint, payload requests.AccountRequest) responses.ServiceResponse {

	// Check portfolio
	if err := service.repoPortfolio.UserPortfolioExists(rctx, payload.PortfolioId, userId); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Portfolio not found!", errors.New("portfolio not found"))
		}
		service.container.Logger.Error("account.appService.create-UserPortfolioExists", "error", err.Error(), "payload", payload, "userId", userId)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong. Please try again later.", err)
	}

	// Check avatar
	if err := service.repoAvatar.AvatarByTypeExists(rctx, payload.AvatarId, commonType.AvatarTypePortfolio); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Avatar not found!", errors.New("avatar not found"))
		}

		service.container.Logger.Error("account.appService.create-AvatarByTypeExists", "error", err.Error(), "avatarId", payload.AvatarId, "avatarType", commonType.AvatarTypePortfolio)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong. Please try again later.", err)
	}

	// Check currencyCode
	if err := service.repoCurrency.CodeExists(rctx, payload.CurrencyCode); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Currency not found!", errors.New("currency not found"))
		}
		service.container.Logger.Error("account.appService.create-CodeExists", "error", err.Error(), "currencyCode", payload.CurrencyCode, "userId", userId)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong. Please try again later.", err)
	}

	id, err := service.repoAccount.Create(rctx, models.Account{
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
		service.container.Logger.Error("account.appService.create-Create", "error", err.Error(), "payload", payload, "userId", userId)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong. Please try again later.", err)
	}

	response, _ := service.repoAccount.Get(rctx, id)
	return responses.SuccessResponse("Account is created!", response)
}

func (service *Account) Update(rctx *requests.RequestContext, id, userId uint, payload requests.AccountUpdateRequest) responses.ServiceResponse {

	// Check avatar
	if err := service.repoAvatar.AvatarByTypeExists(rctx, payload.AvatarId, commonType.AvatarTypePortfolio); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Avatar not found!", errors.New("avatar not found"))
		}
		service.container.Logger.Error("account.appService.update-AvatarByTypeExists", "error", err.Error(), "avatarId", payload.AvatarId, "avatarType", commonType.AvatarTypePortfolio)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong. Please try again later.", err)
	}

	// Check currencyCode
	if err := service.repoCurrency.CodeExists(rctx, payload.CurrencyCode); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Currency not found!", errors.New("currency not found"))
		}
		service.container.Logger.Error("account.appService.update-CodeExists", "error", err.Error(), "currencyCode", payload.CurrencyCode, "userId", userId)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong. Please try again later.", err)
	}

	if err := service.repoAccount.Update(rctx, id, userId, payload); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Account not updated!", errors.New("account not updated"))
		}
		service.container.Logger.Error("account.appService.update-Update", "error", err.Error(), "id", id, "userId", userId, "id", payload)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong. Please try again later.", err)
	}

	response, _ := service.repoAccount.Get(rctx, id)
	return responses.SuccessResponse("Account record updated!", response)
}

func (service *Account) Delete(rctx *requests.RequestContext, id, userId uint) responses.ServiceResponse {
	if err := service.repoAccount.Delete(rctx, id, userId); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Account is not deleted!", errors.New("account not deleted"))
		}
		service.container.Logger.Error("account.appService.delete-Delete", "error", err.Error(), "id", id, "userId", userId)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong. Please try again later.", err)
	}

	return responses.SuccessResponse("Account is deleted!", nil)
}

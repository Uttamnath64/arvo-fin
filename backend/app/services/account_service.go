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
			return responses.ErrorResponse(common.StatusNotFound, "We couldn’t find the account you were looking for.", err)
		}

		service.container.Logger.Error("account.appService.get-Get", "error", err.Error(), "id", id)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	// Response
	return responses.SuccessResponse("Account details retrieved successfully.", response)
}

func (service *Account) GetList(rctx *requests.RequestContext, portfolioId, userId uint) responses.ServiceResponse {
	response, err := service.repoAccount.GetList(rctx, portfolioId, userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "No accounts were found in this portfolio.", err)
		}

		service.container.Logger.Error("account.appService.getList-GetList", "error", err.Error(), "portfolioId", portfolioId, "userId", userId)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	// Response
	return responses.SuccessResponse("Accounts details retrieved successfully.", response)
}

func (service *Account) AccountTypes(rctx *requests.RequestContext) responses.ServiceResponse {

	accountTypes := []responses.AccountTypeResponse{
		{
			Type: commonType.AccountTypeBank,
			Name: commonType.AccountTypeBank.String(),
			Icon: service.repoAvatar.GetByNameAndType(rctx, "Bank", commonType.AvatarTypePortfolio),
		},
		{
			Type: commonType.AccountTypeCash,
			Name: commonType.AccountTypeCash.String(),
			Icon: service.repoAvatar.GetByNameAndType(rctx, "Cash", commonType.AvatarTypePortfolio),
		},
		{
			Type: commonType.AccountTypeWallet,
			Name: commonType.AccountTypeWallet.String(),
			Icon: service.repoAvatar.GetByNameAndType(rctx, "Wallet", commonType.AvatarTypePortfolio),
		},
		{
			Type: commonType.AccountTypeCredit,
			Name: commonType.AccountTypeCredit.String(),
			Icon: service.repoAvatar.GetByNameAndType(rctx, "Credit", commonType.AvatarTypePortfolio),
		},
		{
			Type: commonType.AccountTypeLoan,
			Name: commonType.AccountTypeLoan.String(),
			Icon: service.repoAvatar.GetByNameAndType(rctx, "Loan", commonType.AvatarTypePortfolio),
		},
		{
			Type: commonType.AccountTypeInvestment,
			Name: commonType.AccountTypeInvestment.String(),
			Icon: service.repoAvatar.GetByNameAndType(rctx, "Chart", commonType.AvatarTypePortfolio),
		},
	}

	// Response
	return responses.SuccessResponse("Available account types retrieved successfully.", accountTypes)
}

func (service *Account) Create(rctx *requests.RequestContext, userId uint, payload requests.AccountRequest) responses.ServiceResponse {

	// Check portfolio
	if err := service.repoPortfolio.UserPortfolioExists(rctx, payload.PortfolioId, userId); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "No portfolio found for this user.", errors.New("portfolio not found"))
		}
		service.container.Logger.Error("account.appService.create-UserPortfolioExists", "error", err.Error(), "payload", payload, "userId", userId)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	// Check avatar
	if err := service.repoAvatar.AvatarByTypeExists(rctx, payload.AvatarId, commonType.AvatarTypePortfolio); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Selected avatar not found. Please choose a valid one.", errors.New("avatar not found"))
		}

		service.container.Logger.Error("account.appService.create-AvatarByTypeExists", "error", err.Error(), "avatarId", payload.AvatarId, "avatarType", commonType.AvatarTypePortfolio)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	// Check currencyCode
	if err := service.repoCurrency.CodeExists(rctx, payload.CurrencyCode); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "The specified currency code is invalid or not supported.", errors.New("currency not found"))
		}
		service.container.Logger.Error("account.appService.create-CodeExists", "error", err.Error(), "currencyCode", payload.CurrencyCode, "userId", userId)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
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
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	response, _ := service.repoAccount.Get(rctx, id)
	return responses.SuccessResponse("Your account has been created successfully. 🎉", response)
}

func (service *Account) Update(rctx *requests.RequestContext, id, userId uint, payload requests.AccountUpdateRequest) responses.ServiceResponse {

	// Check avatar
	if err := service.repoAvatar.AvatarByTypeExists(rctx, payload.AvatarId, commonType.AvatarTypePortfolio); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Selected avatar not found. Please choose a valid one.", errors.New("avatar not found"))
		}
		service.container.Logger.Error("account.appService.update-AvatarByTypeExists", "error", err.Error(), "avatarId", payload.AvatarId, "avatarType", commonType.AvatarTypePortfolio)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	// Check currencyCode
	if err := service.repoCurrency.CodeExists(rctx, payload.CurrencyCode); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "The specified currency code is invalid or not supported.", errors.New("currency not found"))
		}
		service.container.Logger.Error("account.appService.update-CodeExists", "error", err.Error(), "currencyCode", payload.CurrencyCode, "userId", userId)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	if err := service.repoAccount.Update(rctx, id, userId, payload); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Unable to update the account. Please try again.", errors.New("account not updated"))
		}
		service.container.Logger.Error("account.appService.update-Update", "error", err.Error(), "id", id, "userId", userId, "id", payload)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	response, _ := service.repoAccount.Get(rctx, id)
	return responses.SuccessResponse("Your account details have been updated successfully.", response)
}

func (service *Account) Delete(rctx *requests.RequestContext, id, userId uint) responses.ServiceResponse {
	if err := service.repoAccount.Delete(rctx, id, userId); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Account deletion failed. Please try again.", errors.New("account not deleted"))
		}
		service.container.Logger.Error("account.appService.delete-Delete", "error", err.Error(), "id", id, "userId", userId)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	return responses.SuccessResponse("The account has been deleted successfully.", nil)
}

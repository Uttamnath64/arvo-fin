package services

import (
	"errors"

	"github.com/Uttamnath64/arvo-fin/app/common"
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/repository"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/responses"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"github.com/Uttamnath64/arvo-fin/pkg/pagination"
	"gorm.io/gorm"
)

type Transaction struct {
	container     *storage.Container
	repo          repository.TransactionRepository
	repoPortfolio repository.PortfolioRepository
	repoAccount   repository.AccountRepository
	repoCategory  repository.CategoryRepository
}

func NewTransaction(container *storage.Container) *Transaction {
	return &Transaction{
		container:     container,
		repo:          repository.NewTransaction(container),
		repoPortfolio: repository.NewPortfolio(container),
		repoAccount:   repository.NewAccount(container),
		repoCategory:  repository.NewCategory(container),
	}
}

func (service *Transaction) Get(rctx *requests.RequestContext, id uint) responses.ServiceResponse {

	response, err := service.repo.Get(rctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "We couldn’t find the transaction you were looking for.", err)
		}

		service.container.Logger.Error("transaction.appService.get-Get", "error", err.Error(), "id", id)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	// Response
	return responses.SuccessResponse("Transaction details retrieved successfully.", response)
}

func (service *Transaction) GetList(rctx *requests.RequestContext, transactionQuery requests.TransactionQuery, pagination pagination.Pagination) responses.ServiceResponse {
	response, err := service.repo.GetList(rctx, transactionQuery, pagination)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Transactions not found.", err)
		}

		service.container.Logger.Error("transaction.appService.getList-GetList", "error", err.Error(), "transactionQuery", transactionQuery, "userId", rctx.UserID, "pagination", pagination)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	// Response
	return responses.SuccessResponse("Transactions details retrieved successfully.", response)
}

func (service *Transaction) Create(rctx *requests.RequestContext, payload requests.TransactionRequest) responses.ServiceResponse {

	// Check portfolio
	if err := service.repoPortfolio.UserPortfolioExists(rctx, payload.PortfolioId, rctx.UserID); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "No portfolio found for this user.", errors.New("portfolio not found"))
		}
		service.container.Logger.Error("transaction.appService.create-UserPortfolioExists", "error", err.Error(), "payload", payload, "user", rctx, "portfolioId", payload.PortfolioId)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	// Check account
	if err := service.repoAccount.UserAccountExists(rctx, payload.AccountId, payload.PortfolioId, rctx.UserID); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "No account found for this protfolio and user.", errors.New("account not found"))
		}
		service.container.Logger.Error("transaction.appService.create-UserAccountExists", "error", err.Error(), "payload", payload, "user", rctx, "portfolioId", payload.PortfolioId, "accountId", payload.AccountId)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	// Check category
	if err := service.repoCategory.UserCategoryExists(rctx, payload.CategoryId, payload.PortfolioId, rctx.UserID); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "No category found for this protfolio and user.", errors.New("category not found"))
		}
		service.container.Logger.Error("transaction.appService.create-UserCategoryExists", "error", err.Error(), "payload", payload, "user", rctx, "portfolioId", payload.PortfolioId, "categoryId", payload.CategoryId)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	// Transaction is transfer account
	if payload.TransferAccountId != nil {
		if err := service.repoAccount.UserAccountExists(rctx, *payload.TransferAccountId, payload.PortfolioId, rctx.UserID); err != nil {
			if err == gorm.ErrRecordNotFound {
				return responses.ErrorResponse(common.StatusNotFound, "No transfer account found for this protfolio and user.", errors.New("transfer account not found"))
			}
			service.container.Logger.Error("transaction.appService.create-UserAccountExists-TransferAccountId", "error", err.Error(), "payload", payload, "user", rctx, "portfolioId", payload.PortfolioId, "transferAccountId", payload.TransferAccountId)
			return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
		}
	}

	data := models.Transaction{
		UserId:            rctx.UserID,
		PortfolioId:       payload.PortfolioId,
		CategoryId:        payload.CategoryId,
		AccountId:         payload.AccountId,
		Amount:            payload.Amount,
		Type:              payload.Type,
		Note:              payload.Note,
		TransferAccountId: payload.TransferAccountId,
	}

	id, err := service.repo.Create(rctx, data)
	if err != nil {
		service.container.Logger.Error("transaction.appService.create-Create", "error", err.Error(), "payload", payload, "user", rctx, "portfolioId", payload.PortfolioId)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	response, _ := service.repo.Get(rctx, id)
	return responses.SuccessResponse("Your transaction has been created successfully. 🎉", response)
}

func (service *Transaction) Update(rctx *requests.RequestContext, id uint, payload requests.TransactionRequest) responses.ServiceResponse {

	// Check portfolio
	if err := service.repoPortfolio.UserPortfolioExists(rctx, payload.PortfolioId, rctx.UserID); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "No portfolio found for this user.", errors.New("portfolio not found"))
		}
		service.container.Logger.Error("transaction.appService.update-UserPortfolioExists", "error", err.Error(), "payload", payload, "user", rctx, "portfolioId", payload.PortfolioId)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	// Check account
	if err := service.repoAccount.UserAccountExists(rctx, payload.AccountId, payload.PortfolioId, rctx.UserID); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "No account found for this protfolio and user.", errors.New("account not found"))
		}
		service.container.Logger.Error("transaction.appService.update-UserAccountExists", "error", err.Error(), "payload", payload, "user", rctx, "portfolioId", payload.PortfolioId, "accountId", payload.AccountId)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	// Check category
	if err := service.repoCategory.UserCategoryExists(rctx, payload.CategoryId, payload.PortfolioId, rctx.UserID); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "No category found for this protfolio and user.", errors.New("category not found"))
		}
		service.container.Logger.Error("transaction.appService.update-UserCategoryExists", "error", err.Error(), "payload", payload, "user", rctx, "portfolioId", payload.PortfolioId, "categoryId", payload.CategoryId)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	// Transaction is transfer account
	if payload.TransferAccountId != nil {
		if err := service.repoAccount.UserAccountExists(rctx, *payload.TransferAccountId, payload.PortfolioId, rctx.UserID); err != nil {
			if err == gorm.ErrRecordNotFound {
				return responses.ErrorResponse(common.StatusNotFound, "No transfer account found for this protfolio and user.", errors.New("transfer account not found"))
			}
			service.container.Logger.Error("transaction.appService.update-UserAccountExists-TransferAccountId", "error", err.Error(), "payload", payload, "user", rctx, "portfolioId", payload.PortfolioId, "transferAccountId", payload.TransferAccountId)
			return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
		}
	}

	if err := service.repo.Update(rctx, id, payload); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Unable to update the account. Please try again.", errors.New("transaction not updated"))
		}
		service.container.Logger.Error("transaction.appService.update-Update", "error", err.Error(), "id", id, "user", rctx, "id", payload)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	response, _ := service.repo.Get(rctx, id)
	return responses.SuccessResponse("Your transaction details have been updated successfully.", response)
}

func (service *Transaction) Delete(rctx *requests.RequestContext, id uint) responses.ServiceResponse {
	if err := service.repo.Delete(rctx, id); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Transaction deletion failed. Please try again.", errors.New("transaction not deleted"))
		}
		service.container.Logger.Error("transaction.appService.delete-Delete", "error", err.Error(), "id", id, "user", rctx)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong on our end. Please try again in a moment.", err)
	}

	return responses.SuccessResponse("The transaction has been deleted successfully.", nil)
}

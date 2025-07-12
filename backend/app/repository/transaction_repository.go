package repository

import (
	"strings"

	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"github.com/Uttamnath64/arvo-fin/pkg/pagination"
	"gorm.io/gorm"
)

type Transaction struct {
	container *storage.Container
}

func NewTransaction(container *storage.Container) *Transaction {
	return &Transaction{
		container: container,
	}
}

func (repo *Transaction) GetList(rctx *requests.RequestContext, transactionQuery requests.TransactionQuery, pagination pagination.Pagination) (*[]models.Transaction, error) {
	var transactions []models.Transaction

	query := repo.container.Config.ReadOnlyDB.WithContext(rctx.Ctx).Preload("Category").Preload("Account").Preload("TransferAccount")

	if rctx.UserType == commonType.UserTypeUser || transactionQuery.UserId > 0 {
		query.Where("user_id = ?", transactionQuery.UserId)
	}
	if transactionQuery.PortfolioId > 0 {
		query.Where("portfolio_id = ?", transactionQuery.PortfolioId)
	}
	if transactionQuery.AccountId > 0 {
		query.Where("account_id = ?", transactionQuery.AccountId)
	}
	if transactionQuery.CategoryId > 0 {
		query.Where("category_id = ?", transactionQuery.CategoryId)
	}
	if !transactionQuery.DateFrom.IsZero() && !transactionQuery.DateTo.IsZero() {
		query.Where("updated_at >= ? AND updated_at <= ?", transactionQuery.DateFrom, transactionQuery.DateTo)
	}
	if strings.TrimSpace(transactionQuery.Search) != "" {
		query.Where("search = ?", transactionQuery.CategoryId)
	}
	if transactionQuery.Type != nil {
		query.Where("type = ?", transactionQuery.Type)
	}
	query.Limit(pagination.Page).Offset(pagination.GetOffset()).
		Order("updated_at " + transactionQuery.Order)

	err := query.Find(&transactions).Error

	if err != nil {
		return nil, err // Other errors
	}
	if len(transactions) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &transactions, nil
}

func (repo *Transaction) Get(rctx *requests.RequestContext, id uint) (*models.Transaction, error) {
	var transaction models.Transaction

	query := repo.container.Config.ReadOnlyDB.WithContext(rctx.Ctx).Preload("Category").Preload("Account").Preload("TransferAccount").Where("id = ?", id)
	if rctx.UserType == commonType.UserTypeUser {
		query.Where("user_id = ?", rctx.UserID)
	}

	if err := query.First(&transaction).Error; err != nil {
		return nil, err // Other errors
	}
	if transaction.ID == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &transaction, nil
}

func (repo *Transaction) Create(rctx *requests.RequestContext, transaction models.Transaction) (uint, error) {
	err := repo.container.Config.ReadWriteDB.WithContext(rctx.Ctx).Create(&transaction).Error
	if err != nil {
		return 0, err
	}
	return transaction.ID, nil
}

func (repo *Transaction) Update(rctx *requests.RequestContext, id uint, payload requests.TransactionRequest) error {
	result := repo.container.Config.ReadWriteDB.WithContext(rctx.Ctx).Model(&models.Account{}).
		Where("id = ? AND user_id = ?", id, rctx.UserID).
		Updates(map[string]interface{}{
			"transfer_account_id": payload.TransferAccountId,
			"account_id":          payload.AccountId,
			"category_id":         payload.CategoryId,
			"amount":              payload.Amount,
			"type":                payload.Type,
			"note":                payload.Note,
		})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (repo *Transaction) Delete(rctx *requests.RequestContext, id uint) error {
	result := repo.container.Config.ReadWriteDB.WithContext(rctx.Ctx).Where("id = ? AND user_id = ?", id, rctx.UserID).Delete(&models.Transaction{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

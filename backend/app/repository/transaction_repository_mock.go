package repository

import (
	"time"

	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"github.com/Uttamnath64/arvo-fin/pkg/pagination"
	"gorm.io/gorm"
)

type TestTransaction struct {
	container *storage.Container
}

func NewTestTransaction(container *storage.Container) *TestTransaction {
	return &TestTransaction{
		container: container,
	}
}

func (repo *TestTransaction) GetList(rctx *requests.RequestContext, transactionQuery requests.TransactionQuery, pagination pagination.Pagination) (*[]models.Transaction, error) {

	transferAccountId := uint(2)
	transactions := []models.Transaction{
		{
			BaseModel: models.BaseModel{
				ID:        3,
				CreatedAt: time.Date(2025, 7, 13, 8, 15, 17, 591000000, time.UTC),
				UpdatedAt: time.Date(2025, 7, 13, 8, 15, 17, 591000000, time.UTC),
			},
			UserId:            2,
			TransferAccountId: &transferAccountId,
			AccountId:         1,
			CategoryId:        1,
			PortfolioId:       3,
			Amount:            100.0,
			Type:              commonType.TransactionType(1),
			Note:              "test",

			Account: &models.Account{
				BaseModel: models.BaseModel{
					ID:        1,
					CreatedAt: time.Date(2025, 7, 13, 7, 28, 9, 644000000, time.UTC),
					UpdatedAt: time.Date(2025, 7, 13, 7, 28, 9, 644000000, time.UTC),
				},
				UserId:         2,
				PortfolioId:    3,
				AvatarId:       47,
				Name:           "CASH name",
				Type:           2,
				CurrencyCode:   "INR",
				OpeningBalance: 10,
				Note:           "testing",
				Avatar:         &models.Avatar{},
				Currency: &models.Currency{
					Code:   "",
					Name:   "",
					Symbol: "",
				},
			},

			TransferAccount: &models.Account{
				BaseModel: models.BaseModel{
					ID:        2,
					CreatedAt: time.Date(2025, 7, 13, 8, 15, 9, 393000000, time.UTC),
					UpdatedAt: time.Date(2025, 7, 13, 8, 15, 9, 393000000, time.UTC),
				},
				UserId:         2,
				PortfolioId:    3,
				AvatarId:       47,
				Name:           "bank",
				Type:           2,
				CurrencyCode:   "INR",
				OpeningBalance: 10,
				Note:           "testing",
				Avatar:         &models.Avatar{},
				Currency:       &models.Currency{},
			},

			Category: &models.Category{
				BaseModel: models.BaseModel{
					ID:        1,
					CreatedAt: time.Date(2025, 7, 6, 14, 3, 4, 154000000, time.UTC),
					UpdatedAt: time.Date(2025, 7, 6, 14, 3, 4, 154000000, time.UTC),
				},
				SourceId:     1,
				AvatarId:     16,
				SourceType:   2,
				PortfolioId:  nil,
				CopiedFromId: nil,
				Name:         "Salary",
				Type:         1,
				Avatar:       &models.Avatar{},
			},

			TransactionAudit: nil,
		},
		{
			BaseModel: models.BaseModel{
				ID:        4,
				CreatedAt: time.Date(2025, 7, 13, 8, 15, 17, 591000000, time.UTC),
				UpdatedAt: time.Date(2025, 7, 13, 8, 15, 17, 591000000, time.UTC),
			},
			UserId:            2,
			TransferAccountId: &transferAccountId,
			AccountId:         1,
			CategoryId:        1,
			PortfolioId:       3,
			Amount:            100.0,
			Type:              commonType.TransactionType(1),
			Note:              "test",

			Account: &models.Account{
				BaseModel: models.BaseModel{
					ID:        1,
					CreatedAt: time.Date(2025, 7, 13, 7, 28, 9, 644000000, time.UTC),
					UpdatedAt: time.Date(2025, 7, 13, 7, 28, 9, 644000000, time.UTC),
				},
				UserId:         2,
				PortfolioId:    3,
				AvatarId:       47,
				Name:           "CASH name",
				Type:           2,
				CurrencyCode:   "INR",
				OpeningBalance: 10,
				Note:           "testing",
				Avatar:         &models.Avatar{},
				Currency: &models.Currency{
					Code:   "",
					Name:   "",
					Symbol: "",
				},
			},

			TransferAccount: &models.Account{
				BaseModel: models.BaseModel{
					ID:        2,
					CreatedAt: time.Date(2025, 7, 13, 8, 15, 9, 393000000, time.UTC),
					UpdatedAt: time.Date(2025, 7, 13, 8, 15, 9, 393000000, time.UTC),
				},
				UserId:         2,
				PortfolioId:    3,
				AvatarId:       47,
				Name:           "bank",
				Type:           2,
				CurrencyCode:   "INR",
				OpeningBalance: 10,
				Note:           "testing",
				Avatar:         &models.Avatar{},
				Currency:       &models.Currency{},
			},

			Category: &models.Category{
				BaseModel: models.BaseModel{
					ID:        1,
					CreatedAt: time.Date(2025, 7, 6, 14, 3, 4, 154000000, time.UTC),
					UpdatedAt: time.Date(2025, 7, 6, 14, 3, 4, 154000000, time.UTC),
				},
				SourceId:     1,
				AvatarId:     16,
				SourceType:   2,
				PortfolioId:  nil,
				CopiedFromId: nil,
				Name:         "Salary",
				Type:         1,
				Avatar:       &models.Avatar{},
			},

			TransactionAudit: nil,
		},
	}
	return &transactions, nil
}

func (repo *TestTransaction) Get(rctx *requests.RequestContext, id uint) (*models.Transaction, error) {
	if id != 1 {
		return nil, gorm.ErrRecordNotFound
	}

	transaction := models.Transaction{
		BaseModel: models.BaseModel{
			ID:        3,
			CreatedAt: time.Date(2025, 7, 13, 8, 15, 17, 591000000, time.UTC),
			UpdatedAt: time.Date(2025, 7, 13, 8, 15, 17, 591000000, time.UTC),
		},
		UserId:      2,
		AccountId:   1,
		CategoryId:  1,
		PortfolioId: 3,
		Amount:      100.0,
		Type:        commonType.TransactionType(1),
		Note:        "test",

		Account: &models.Account{
			BaseModel: models.BaseModel{
				ID:        1,
				CreatedAt: time.Date(2025, 7, 13, 7, 28, 9, 644000000, time.UTC),
				UpdatedAt: time.Date(2025, 7, 13, 7, 28, 9, 644000000, time.UTC),
			},
			UserId:         2,
			PortfolioId:    3,
			AvatarId:       47,
			Name:           "CASH name",
			Type:           2,
			CurrencyCode:   "INR",
			OpeningBalance: 10,
			Note:           "testing",
			Avatar:         &models.Avatar{},
			Currency: &models.Currency{
				Code:   "",
				Name:   "",
				Symbol: "",
			},
		},

		Category: &models.Category{
			BaseModel: models.BaseModel{
				ID:        1,
				CreatedAt: time.Date(2025, 7, 6, 14, 3, 4, 154000000, time.UTC),
				UpdatedAt: time.Date(2025, 7, 6, 14, 3, 4, 154000000, time.UTC),
			},
			SourceId:     1,
			AvatarId:     16,
			SourceType:   2,
			PortfolioId:  nil,
			CopiedFromId: nil,
			Name:         "Salary",
			Type:         1,
			Avatar:       &models.Avatar{},
		},

		TransactionAudit: nil,
	}
	return &transaction, nil
}

func (repo *TestTransaction) Create(rctx *requests.RequestContext, transaction models.Transaction) (uint, error) {
	return 1, nil
}

func (repo *TestTransaction) Update(rctx *requests.RequestContext, id uint, payload requests.TransactionRequest) error {
	if id == 1 {
		return nil
	}
	return gorm.ErrRecordNotFound
}

func (repo *TestTransaction) Delete(rctx *requests.RequestContext, id uint) error {
	if id == 1 {
		return nil
	}
	return gorm.ErrRecordNotFound
}

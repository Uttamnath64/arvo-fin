package repository

import (
	"time"

	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"gorm.io/gorm"
)

type TestCategory struct {
	container *storage.Container
}

func NewTestCategory(container *storage.Container) *TestCategory {
	return &TestCategory{
		container: container,
	}
}

func (repo *TestCategory) GetList(rctx *requests.RequestContext, portfolioId, userId uint) (*[]models.Category, error) {
	if userId != 1 || portfolioId != 1 {
		return nil, gorm.ErrRecordNotFound
	}
	copiedFromId := uint(23)
	return &[]models.Category{
		{
			BaseModel: models.BaseModel{
				ID:        1,
				CreatedAt: time.Now().Add(-1 * time.Hour),
				UpdatedAt: time.Now(),
			},
			SourceId:     1,
			SourceType:   commonType.UserTypeUser,
			PortfolioId:  &portfolioId,
			Name:         "Test Account",
			Type:         commonType.TransactionTypeExpense,
			CopiedFromId: &copiedFromId,
			Avatar: models.Avatar{
				BaseModel: models.BaseModel{
					ID:        1,
					CreatedAt: time.Now().Add(-1 * time.Hour),
					UpdatedAt: time.Now(),
				},
				Name: "Test Avatar",
				Icon: "T",
				Type: commonType.AvatarTypeDefault,
			},
		},
		{
			BaseModel: models.BaseModel{
				ID:        2,
				CreatedAt: time.Now().Add(-1 * time.Hour),
				UpdatedAt: time.Now(),
			},
			SourceId:    1,
			SourceType:  commonType.UserTypeUser,
			PortfolioId: &portfolioId,
			Name:        "Test Account",
			Type:        commonType.TransactionTypeExpense,
			Avatar: models.Avatar{
				BaseModel: models.BaseModel{
					ID:        1,
					CreatedAt: time.Now().Add(-1 * time.Hour),
					UpdatedAt: time.Now(),
				},
				Name: "Test Avatar 2",
				Icon: "G",
				Type: commonType.AvatarTypeDefault,
			},
		},
	}, nil
}

func (repo *TestCategory) Get(rctx *requests.RequestContext, id, userId uint) (*models.Category, error) {
	if id != 1 {
		return nil, gorm.ErrRecordNotFound
	}
	portfolioId := uint(12)
	return &models.Category{
		BaseModel: models.BaseModel{
			ID:        1,
			CreatedAt: time.Now().Add(-1 * time.Hour),
			UpdatedAt: time.Now(),
		},
		SourceId:    1,
		SourceType:  commonType.UserTypeUser,
		PortfolioId: &portfolioId,
		Name:        "Test Account",
		Type:        commonType.TransactionTypeExpense,
		Avatar: models.Avatar{
			BaseModel: models.BaseModel{
				ID:        1,
				CreatedAt: time.Now().Add(-1 * time.Hour),
				UpdatedAt: time.Now(),
			},
			Name: "Test Avatar 2",
			Icon: "G",
			Type: commonType.AvatarTypeDefault,
		},
	}, nil
}

func (repo *TestCategory) Create(rctx *requests.RequestContext, category models.Category) (uint, error) {
	return 1, nil
}

func (repo *TestCategory) Update(rctx *requests.RequestContext, id, userId uint, payload requests.CategoryRequest) error {
	if id != 1 || userId != 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (repo *TestCategory) Delete(rctx *requests.RequestContext, id, userId uint) error {
	if id != 1 || userId != 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

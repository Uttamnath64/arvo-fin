package repository

import (
	"time"

	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"gorm.io/gorm"
)

type TestPortfolio struct {
	container *storage.Container
}

func NewTestPortfolio(container *storage.Container) *TestPortfolio {
	return &TestPortfolio{
		container: container,
	}
}

func (repo *TestPortfolio) UserPortfolioExists(rctx *requests.RequestContext, id, userId uint) error {
	if userId == 1 {
		return nil
	}
	return gorm.ErrRecordNotFound
}

func (repo *TestPortfolio) GetList(rctx *requests.RequestContext, userId uint) (*[]models.Portfolio, error) {
	if userId == 1 {
		portfolios := []models.Portfolio{
			{
				BaseModel: models.BaseModel{
					ID:        1,
					CreatedAt: time.Now().Add(-1 * time.Hour),
					UpdatedAt: time.Now(),
				},
				UserId:   1,
				AvatarId: 1,
				Name:     "Retirement Fund",
				Avatar: models.Avatar{
					BaseModel: models.BaseModel{
						ID:        2,
						CreatedAt: time.Now().Add(-1 * time.Hour),
						UpdatedAt: time.Now(),
					},
					Name: "Crypto Portfolio",
					Icon: "T",
					Type: commonType.AvatarTypeDefault,
				},
			},
			{
				BaseModel: models.BaseModel{
					ID:        1,
					CreatedAt: time.Now().Add(-1 * time.Hour),
					UpdatedAt: time.Now(),
				},
				UserId:   1,
				AvatarId: 1,
				Name:     "Retirement Fund",
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
		}
		return &portfolios, nil
	}
	return nil, gorm.ErrRecordNotFound
}

func (repo *TestPortfolio) Get(rctx *requests.RequestContext, id, userId uint, userType commonType.UserType) (*models.Portfolio, error) {
	if id == 1 && userId == 1 && userType == commonType.UserTypeUser {
		portfolios := models.Portfolio{
			BaseModel: models.BaseModel{
				ID:        1,
				CreatedAt: time.Now().Add(-1 * time.Hour),
				UpdatedAt: time.Now(),
			},
			UserId:   1,
			AvatarId: 1,
			Name:     "Retirement Fund",
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
		}
		return &portfolios, nil
	}
	return nil, gorm.ErrRecordNotFound
}

func (repo *TestPortfolio) Create(rctx *requests.RequestContext, portfolio models.Portfolio) error {
	return nil
}

func (repo *TestPortfolio) Update(rctx *requests.RequestContext, id, userId uint, payload requests.PortfolioRequest) error {
	if id == 1 && userId == 1 {
		return nil
	}
	return gorm.ErrRecordNotFound
}

func (repo *TestPortfolio) Delete(rctx *requests.RequestContext, id, userId uint) error {
	if id == 1 && userId == 1 {
		return nil
	}
	return gorm.ErrRecordNotFound
}

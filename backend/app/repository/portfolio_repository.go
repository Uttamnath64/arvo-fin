package repository

import (
	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"gorm.io/gorm"
)

type Portfolio struct {
	container *storage.Container
}

func NewPortfolio(container *storage.Container) *Portfolio {
	return &Portfolio{
		container: container,
	}
}

func (repo *Portfolio) UserPortfolioExists(rctx *requests.RequestContext, id, userId uint) error {
	var count int64

	err := repo.container.Config.ReadOnlyDB.WithContext(rctx.Ctx).Model(&models.Portfolio{}).
		Where("id = ? and user_id = ?", id, userId).Count(&count).Error

	if err != nil {
		return err
	}
	if count == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (repo *Portfolio) GetList(rctx *requests.RequestContext, userId uint) (*[]models.Portfolio, error) {
	var portfolios []models.Portfolio

	err := repo.container.Config.ReadOnlyDB.WithContext(rctx.Ctx).Preload("Avatar").Where("user_id = ? ", userId).
		Find(&portfolios).Error

	if err != nil {
		return nil, err // Other errors
	}
	if len(portfolios) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &portfolios, nil
}

func (repo *Portfolio) Get(rctx *requests.RequestContext, id, userId uint, userType commonType.UserType) (*models.Portfolio, error) {
	var portfolio models.Portfolio

	query := repo.container.Config.ReadOnlyDB.WithContext(rctx.Ctx).Preload("Avatar").Where("id = ?", id)
	if userType == commonType.UserTypeUser {
		query = query.Where("user_id = ?", userId)
	}
	err := query.First(&portfolio).Error

	if err != nil {
		return nil, err // Other errors
	}
	if portfolio.ID == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &portfolio, nil
}

func (repo *Portfolio) Create(rctx *requests.RequestContext, portfolio models.Portfolio) error {
	return repo.container.Config.ReadWriteDB.WithContext(rctx.Ctx).Create(&portfolio).Error
}

func (repo *Portfolio) Update(rctx *requests.RequestContext, id, userId uint, payload requests.PortfolioRequest) error {
	result := repo.container.Config.ReadWriteDB.WithContext(rctx.Ctx).Model(&models.Portfolio{}).
		Where("id = ? AND user_id = ?", id, userId).
		Updates(map[string]interface{}{
			"name":      payload.Name,
			"avatar_id": payload.AvatarId,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (repo *Portfolio) Delete(rctx *requests.RequestContext, id, userId uint) error {
	result := repo.container.Config.ReadWriteDB.WithContext(rctx.Ctx).Where("id = ? AND user_id = ?", id, userId).Delete(&models.Portfolio{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

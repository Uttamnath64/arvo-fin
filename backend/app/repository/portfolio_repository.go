package repository

import (
	"errors"

	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/responses"
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

func (repo *Portfolio) GetList(userId uint, userType commonType.UserType) (*[]responses.PortfolioResponse, error) {
	var portfolio models.Portfolio
	var avatar models.Avatar

	var portfolios []responses.PortfolioResponse

	query := repo.container.Config.ReadOnlyDB.Table(portfolio.GetName() + " p").
		Joins("JOIN " + avatar.GetName() + " a ON a.id = p.avatar_id").Where("p.deleted_at IS NULL")
	if userType == commonType.User {
		query = query.Where("p.user_id = ? ", userId)
	}
	err := query.Select("p.id, p.name, a.id as avatar_id, a.url").
		Scan(&portfolios).Error

	if err != nil {
		return nil, err // Other errors
	}
	if len(portfolios) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &portfolios, nil
}

func (repo *Portfolio) Get(id, userId uint, userType commonType.UserType) (*responses.PortfolioResponse, error) {
	var portfolio models.Portfolio
	var avatar models.Avatar

	var portfolios responses.PortfolioResponse

	query := repo.container.Config.ReadOnlyDB.Table(portfolio.GetName()+" p").
		Joins("JOIN "+avatar.GetName()+" a ON a.id = p.avatar_id").Where("p.id = ? AND p.deleted_at IS NULL", id)
	if userType == commonType.User {
		query = query.Where("p.user_id = ?", userId)
	}
	err := query.Select("p.id, p.name, a.id as avatar_id, a.url").
		Scan(&portfolios).Error

	if err != nil {
		return nil, err // Other errors
	}
	if portfolios.Id == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &portfolios, nil
}

func (repo *Portfolio) Add(portfolio models.Portfolio) error {
	return repo.container.Config.ReadWriteDB.Create(&portfolio).Error
}

func (repo *Portfolio) Update(id, userId uint, payload requests.PortfolioRequest) error {
	result := repo.container.Config.ReadWriteDB.Model(&models.Portfolio{}).
		Where("id = ? AND user_id = ?", id, userId).
		Updates(map[string]interface{}{
			"name":      payload.Name,
			"avatar_id": payload.AvatarId,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("Portfolio not found!")
	}
	return nil
}

func (repo *Portfolio) Delete(id, userId uint) error {
	result := repo.container.Config.ReadWriteDB.Where("id = ? AND user_id = ?", id, userId).Delete(&models.Portfolio{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("Portfolio not found!")
	}
	return nil
}

package repository

import (
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"gorm.io/gorm"
)

type Currency struct {
	container *storage.Container
}

func NewCurrency(container *storage.Container) *Currency {
	return &Currency{
		container: container,
	}
}

func (repo *Currency) CodeExists(rctx *requests.RequestContext, code string) error {
	var count int64

	err := repo.container.Config.ReadOnlyDB.WithContext(rctx.Ctx).Model(&models.Currency{}).
		Where("code = ?", code).Count(&count).Error

	if err != nil {
		return err
	}
	if count == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

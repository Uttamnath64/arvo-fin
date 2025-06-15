package repository

import (
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/storage"
)

type Currency struct {
	container *storage.Container
}

func NewCurrency(container *storage.Container) *Currency {
	return &Currency{
		container: container,
	}
}

func (repo *Currency) CodeExists(code string) (bool, error) {
	var count int64

	err := repo.container.Config.ReadOnlyDB.Model(&models.Currency{}).
		Where("code = ?", code).Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

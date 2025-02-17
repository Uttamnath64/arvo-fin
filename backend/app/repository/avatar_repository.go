package repository

import (
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/storage"
)

type Avatar struct {
	container *storage.Container
}

func NewAvatar(container *storage.Container) *Avatar {
	return &Avatar{
		container: container,
	}
}

func (repo *Avatar) GetAvatar(id uint, avatar *models.Avatar) error {
	if err := repo.container.Config.ReadOnlyDB.Where("id = ?", id).First(avatar).Error; err != nil {
		return err
	}
	return nil
}

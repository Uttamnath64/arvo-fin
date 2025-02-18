package repository

import (
	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
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

func (repo *Avatar) GetAvatarByType(id uint, avatarType commonType.AvatarType, avatar *models.Avatar) error {
	if err := repo.container.Config.ReadOnlyDB.Where("id = ? AND type = ?", id, avatarType).First(avatar).Error; err != nil {
		return err
	}
	return nil
}

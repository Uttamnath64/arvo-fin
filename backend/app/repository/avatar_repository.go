package repository

import (
	"errors"

	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"gorm.io/gorm"
)

type Avatar struct {
	container *storage.Container
}

func NewAvatar(container *storage.Container) *Avatar {
	return &Avatar{
		container: container,
	}
}

func (repo *Avatar) Get(id uint, avatar *models.Avatar) error {
	return repo.container.Config.ReadOnlyDB.Where("id = ?", id).First(avatar).Error
}

func (repo *Avatar) AvatarByTypeExists(id uint, avatarType commonType.AvatarType) (bool, error) {
	var count int64

	err := repo.container.Config.ReadOnlyDB.Model(&models.User{}).
		Where("id = ? AND type = ?", id, avatarType).Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (repo *Avatar) GetAvatarsByType(avatarType commonType.AvatarType) (*[]models.Avatar, error) {
	var response []models.Avatar
	if err := repo.container.Config.ReadOnlyDB.Model(&models.Avatar{}).Where("type = ?", avatarType).Scan(&response).Error; err != nil {
		return nil, err
	}

	if len(response) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &response, nil
}

func (repo *Avatar) Create(avatar models.Avatar) error {
	return repo.container.Config.ReadWriteDB.Create(&avatar).Error
}

func (repo *Avatar) Update(id uint, payload requests.AvatarRequest) error {
	result := repo.container.Config.ReadWriteDB.Model(&models.Portfolio{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"name": payload.Name,
			"icon": payload.Icon,
			"type": payload.Type,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("Avatar not found!")
	}
	return nil
}

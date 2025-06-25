package repository

import (
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

func (repo *Avatar) Get(id uint) (*models.Avatar, error) {
	var avatar models.Avatar
	return &avatar, repo.container.Config.ReadOnlyDB.Where("id = ?", id).First(&avatar).Error
}

func (repo *Avatar) GetByNameAndType(name string, avatarType commonType.AvatarType) *models.Avatar {
	var avatar models.Avatar
	repo.container.Config.ReadOnlyDB.Where("name = ? and type = ?", name, avatarType).First(&avatar)
	return &avatar
}

func (repo *Avatar) AvatarByTypeExists(id uint, avatarType commonType.AvatarType) error {
	var count int64

	err := repo.container.Config.ReadOnlyDB.Model(&models.Avatar{}).
		Where("id = ? AND type = ?", id, avatarType).Count(&count).Error

	if err != nil {
		return err
	}
	if count == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
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

func (repo *Avatar) Create(avatar models.Avatar) (uint, error) {
	return avatar.ID, repo.container.Config.ReadWriteDB.Create(&avatar).Error
}

func (repo *Avatar) Update(id uint, payload requests.AvatarRequest) error {
	result := repo.container.Config.ReadWriteDB.Model(&models.Avatar{}).
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
		return gorm.ErrRecordNotFound
	}
	return nil
}

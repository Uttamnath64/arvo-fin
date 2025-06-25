package repository

import (
	"time"

	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"gorm.io/gorm"
)

type TestAvatar struct {
	container *storage.Container
}

func NewTestAvatar(container *storage.Container) *TestAvatar {
	return &TestAvatar{
		container: container,
	}
}

func (repo *TestAvatar) Get(id uint) (*models.Avatar, error) {
	if id == 1 {
		return &models.Avatar{
			BaseModel: models.BaseModel{
				ID:        1,
				CreatedAt: time.Now().Add(-2 * time.Hour),
				UpdatedAt: time.Now(),
			},
			Name: "Default Avatar",
			Icon: "🧠",
			Type: commonType.AvatarTypeUser,
		}, nil
	}
	return nil, gorm.ErrRecordNotFound
}

func (repo *TestAvatar) GetByNameAndType(name string, avatarType commonType.AvatarType) *models.Avatar {
	return &models.Avatar{
		Name: name,
		Type: avatarType,
		Icon: "T",
	}
}

func (repo *TestAvatar) AvatarByTypeExists(id uint, avatarType commonType.AvatarType) error {
	if id == 1 {
		return nil
	}
	return gorm.ErrRecordNotFound
}

func (repo *TestAvatar) GetAvatarsByType(avatarType commonType.AvatarType) (*[]models.Avatar, error) {
	responses := []models.Avatar{
		{
			BaseModel: models.BaseModel{
				ID:        1,
				CreatedAt: time.Now().Add(-2 * time.Hour),
				UpdatedAt: time.Now(),
			},
			Name: "Avatar 1",
			Icon: "🧠",
			Type: commonType.AvatarTypeUser,
		},
		{
			BaseModel: models.BaseModel{
				ID:        1,
				CreatedAt: time.Now().Add(-2 * time.Hour),
				UpdatedAt: time.Now(),
			},
			Name: "Avatar 2",
			Icon: "🧠",
			Type: commonType.AvatarTypeUser,
		},
	}
	return &responses, nil
}

func (repo *TestAvatar) Create(payload models.Avatar) (uint, error) {
	return 1, nil
}

func (repo *TestAvatar) Update(id uint, payload requests.AvatarRequest) error {
	if id != 1 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

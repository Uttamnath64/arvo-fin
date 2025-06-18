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

func (repo *TestAvatar) Get(id uint, avatar *models.Avatar) error {
	if id == 1 {
		*avatar = models.Avatar{
			Model: gorm.Model{
				ID:        1,
				CreatedAt: time.Now().Add(-2 * time.Hour),
				UpdatedAt: time.Now(),
			},
			Name: "Default Avatar",
			Icon: "🧠",
			Type: commonType.AvatarTypeUser,
		}
		return nil
	}
	return gorm.ErrRecordNotFound
}

func (repo *TestAvatar) AvatarByTypeExists(id uint, avatarType commonType.AvatarType) (bool, error) {
	if id == 1 {
		return true, nil
	}
	return false, gorm.ErrRecordNotFound
}

func (repo *TestAvatar) GetAvatarsByType(avatarType commonType.AvatarType) (*[]models.Avatar, error) {
	responses := []models.Avatar{
		{
			Model: gorm.Model{
				ID:        1,
				CreatedAt: time.Now().Add(-2 * time.Hour),
				UpdatedAt: time.Now(),
			},
			Name: "Avatar 1",
			Icon: "🧠",
			Type: commonType.AvatarTypeUser,
		},
		{
			Model: gorm.Model{
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

func (repo *TestAvatar) Create(payload models.Avatar) error {
	return nil
}

func (repo *TestAvatar) Update(id uint, payload requests.AvatarRequest) error {
	if id != 1 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

package repository

import (
	"time"

	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/models"
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

func (repo *TestAvatar) GetAvatar(id uint, avatar *models.Avatar) error {
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

func (repo *TestAvatar) GetAvatarByType(id uint, avatarType commonType.AvatarType, avatar *models.Avatar) error {
	if id == 1 {
		*avatar = models.Avatar{
			Model: gorm.Model{
				ID:        1,
				CreatedAt: time.Now().Add(-2 * time.Hour),
				UpdatedAt: time.Now(),
			},
			Name: "Default Avatar",
			Icon: "🧠",
			Type: avatarType,
		}
		return nil
	}
	return gorm.ErrRecordNotFound
}

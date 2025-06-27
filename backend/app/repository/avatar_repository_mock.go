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

func (repo *TestAvatar) Get(rctx *requests.RequestContext, id uint) (*models.Avatar, error) {
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

func (repo *TestAvatar) GetByNameAndType(rctx *requests.RequestContext, name string, avatarType commonType.AvatarType) *models.Avatar {
	return &models.Avatar{
		Name: name,
		Type: avatarType,
		Icon: "T",
	}
}

func (repo *TestAvatar) AvatarByTypeExists(rctx *requests.RequestContext, id uint, avatarType commonType.AvatarType) error {
	if id == 1 {
		return nil
	}
	return gorm.ErrRecordNotFound
}

func (repo *TestAvatar) GetAvatarsByType(rctx *requests.RequestContext, avatarType commonType.AvatarType) (*[]models.Avatar, error) {
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

func (repo *TestAvatar) Create(rctx *requests.RequestContext, payload models.Avatar) (uint, error) {
	return 1, nil
}

func (repo *TestAvatar) Update(rctx *requests.RequestContext, id uint, payload requests.AvatarRequest) error {
	if id != 1 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

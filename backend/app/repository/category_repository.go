package repository

import (
	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"gorm.io/gorm"
)

type Category struct {
	container *storage.Container
}

func NewCategory(container *storage.Container) *Category {
	return &Category{
		container: container,
	}
}

func (repo *Category) GetList(rctx *requests.RequestContext, portfolioId, userId uint) (*[]models.Category, error) {
	var categories []models.Category

	query := `
		SELECT c.id
		FROM categories c
		WHERE c.source_type = ? AND c.source_id = ?

		UNION

		SELECT admin.id
		FROM categories admin
		LEFT JOIN categories user_copy
			ON user_copy.copied_from_id = admin.id
				AND user_copy.source_type = ?
				AND user_copy.source_id = ?
		WHERE admin.source_type = ?
			AND user_copy.id IS NULL`

	var categoryIDs []uint
	err := repo.container.Config.ReadOnlyDB.WithContext(rctx.Ctx).Raw(query,
		commonType.UserTypeUser, userId, // For user-owned categories
		commonType.UserTypeUser, userId, // For checking copied
		commonType.UserTypeAdmin, // Admin categories
	).Pluck("id", &categoryIDs).Error
	if err != nil {
		return nil, err // Other errors
	}
	if len(categoryIDs) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	err = repo.container.Config.ReadOnlyDB.Preload("Avatar").Find(&categories, categoryIDs).Error
	if err != nil {
		return nil, err // Other errors
	}
	return &categories, nil
}

func (repo *Category) Get(rctx *requests.RequestContext, id, userId uint) (*models.Category, error) {
	var category models.Category

	if err := repo.container.Config.ReadOnlyDB.WithContext(rctx.Ctx).Preload("Avatar").
		Where("id = ? AND (source_type = ? OR (source_type = ? AND source_id = ?))",
			id,
			commonType.UserTypeAdmin,
			commonType.UserTypeUser,
			userId,
		).
		First(&category).Error; err != nil {
		return nil, err // Other errors
	}
	if category.AvatarId == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &category, nil
}

func (repo *Category) Create(rctx *requests.RequestContext, category models.Category) (uint, error) {
	err := repo.container.Config.ReadWriteDB.WithContext(rctx.Ctx).Create(&category).Error
	if err != nil {
		return 0, err
	}
	return category.ID, nil
}

func (repo *Category) Update(rctx *requests.RequestContext, id, userId uint, payload requests.CategoryRequest) error {
	result := repo.container.Config.ReadWriteDB.WithContext(rctx.Ctx).Model(&models.Category{}).
		Where("id = ? AND source_type = ? AND source_id = ?", id, commonType.UserTypeUser, userId).
		Updates(map[string]interface{}{
			"name":      payload.Name,
			"avatar_id": payload.AvatarId,
			"type":      payload.Type,
		})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (repo *Category) Delete(rctx *requests.RequestContext, id, userId uint) error {
	result := repo.container.Config.ReadWriteDB.WithContext(rctx.Ctx).
		Where("id = ? AND source_type = ? AND source_id = ?", id, commonType.UserTypeUser, userId).
		Delete(&models.Category{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

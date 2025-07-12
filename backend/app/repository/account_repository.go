package repository

import (
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"gorm.io/gorm"
)

type Account struct {
	container *storage.Container
}

func NewAccount(container *storage.Container) *Account {
	return &Account{
		container: container,
	}
}

func (repo *Account) GetList(rctx *requests.RequestContext, portfolioId, userId uint) (*[]models.Account, error) {
	var account []models.Account

	if err := repo.container.Config.ReadOnlyDB.WithContext(rctx.Ctx).Preload("Avatar").Preload("Currency").Where("user_id = ? AND portfolio_id = ?", userId, portfolioId).
		Find(&account).Error; err != nil {
		return nil, err // Other errors
	}
	if len(account) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &account, nil
}

func (repo *Account) Get(rctx *requests.RequestContext, id uint) (*models.Account, error) {
	var account models.Account

	if err := repo.container.Config.ReadOnlyDB.WithContext(rctx.Ctx).Preload("Avatar").Preload("Currency").Where("id = ?", id).
		First(&account).Error; err != nil {
		return nil, err // Other errors
	}
	if account.AvatarId == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &account, nil
}

func (repo *Account) Create(rctx *requests.RequestContext, account models.Account) (uint, error) {
	err := repo.container.Config.ReadWriteDB.WithContext(rctx.Ctx).Create(&account).Error
	if err != nil {
		return 0, err
	}
	return account.ID, nil
}

func (repo *Account) Update(rctx *requests.RequestContext, id, userId uint, payload requests.AccountUpdateRequest) error {
	result := repo.container.Config.ReadWriteDB.WithContext(rctx.Ctx).Model(&models.Account{}).
		Where("id = ? AND user_id = ?", id, userId).
		Updates(map[string]interface{}{
			"name":          payload.Name,
			"avatar_id":     payload.AvatarId,
			"type":          payload.Type,
			"currency_code": payload.CurrencyCode,
			"note":          payload.Note,
		})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (repo *Account) Delete(rctx *requests.RequestContext, id, userId uint) error {
	result := repo.container.Config.ReadWriteDB.WithContext(rctx.Ctx).Where("id = ? AND user_id = ?", id, userId).Delete(&models.Account{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (repo *Account) UserAccountExists(rctx *requests.RequestContext, id, portfolioId, userId uint) error {
	var count int64

	err := repo.container.Config.ReadOnlyDB.WithContext(rctx.Ctx).Model(&models.Account{}).
		Where("id = ? AND portfolio_id = ? AND user_id = ?", id, portfolioId, userId).Count(&count).Error

	if err != nil {
		return err
	}
	if count == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

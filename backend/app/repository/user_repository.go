package repository

import (
	"strings"

	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/responses"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"gorm.io/gorm"
)

type User struct {
	container *storage.Container
}

func NewUser(container *storage.Container) *User {
	return &User{
		container: container,
	}
}

func (repo *User) GetUserByUsernameOrEmail(rctx *requests.RequestContext, username string, email string, user *models.User) error {
	return repo.container.Config.ReadOnlyDB.WithContext(rctx.Ctx).Model(&models.User{}).
		Where("username = ? or email = ?", username, strings.ToLower(email)).First(user).Error
}

func (repo *User) UsernameExists(rctx *requests.RequestContext, username string) error {
	var count int64

	err := repo.container.Config.ReadOnlyDB.WithContext(rctx.Ctx).Model(&models.User{}).
		Where("username = ?", username).Count(&count).Error

	if err != nil {
		return err
	}
	if count == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (repo *User) EmailExists(rctx *requests.RequestContext, email string) error {
	var count int64

	err := repo.container.Config.ReadOnlyDB.WithContext(rctx.Ctx).Model(&models.User{}).
		Where("email = ?", strings.ToLower(email)).Count(&count).Error

	if err != nil {
		return err
	}
	if count == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (repo *User) CreateUser(rctx *requests.RequestContext, user *models.User) (uint, error) {
	err := repo.container.Config.ReadWriteDB.WithContext(rctx.Ctx).Create(user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (repo *User) UpdatePasswordByEmail(rctx *requests.RequestContext, email, newPassword string) error {
	result := repo.container.Config.ReadWriteDB.WithContext(rctx.Ctx).Model(&models.User{}).
		Where("email = ?", email).
		Update("password", newPassword)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (repo *User) GetUser(rctx *requests.RequestContext, userId uint, user *models.User) error {
	if err := repo.container.Config.ReadOnlyDB.WithContext(rctx.Ctx).Where("id = ?", userId).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *User) Get(rctx *requests.RequestContext, userId uint) (*responses.MeResponse, error) {
	var user models.User
	var avatar models.Avatar
	var response responses.MeResponse

	query := repo.container.Config.ReadOnlyDB.WithContext(rctx.Ctx).Table(user.GetName()+" u").
		Joins("JOIN "+avatar.GetName()+" a ON a.id = u.avatar_id").Where("u.id = ?", userId)

	err := query.Select("u.id, u.name, u.username, u.email, a.id as avatar_id, a.icon as avatar_icon").
		Scan(&response).Error

	if err != nil {
		return nil, err // Other errors
	}
	if response.Id == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &response, nil
}

func (repo *User) GetSettings(rctx *requests.RequestContext, userId uint) (*responses.SettingsResponse, error) {

	var user models.User
	var currency models.Currency
	var response responses.SettingsResponse

	err := repo.container.Config.ReadOnlyDB.WithContext(rctx.Ctx).Table(user.GetName()+" u").
		Joins("JOIN "+currency.GetName()+" c ON c.code = u.currency_code").Where("u.id = ?", userId).
		Select("u.id, u.decimal_places, u.number_format, u.email_notifications, c.code AS 'currency_code', c.symbol AS 'currency_symbol', c.name AS 'currency_name'").
		Scan(&response).Error

	if err != nil {
		return nil, err // Other errors
	}
	if response.Id == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &response, nil
}

func (repo *User) Update(rctx *requests.RequestContext, userId uint, payload requests.MeRequest) error {
	result := repo.container.Config.ReadWriteDB.WithContext(rctx.Ctx).Model(&models.User{}).
		Where("id = ?", userId).
		Updates(map[string]interface{}{
			"name":      payload.Name,
			"username":  payload.Username,
			"avatar_id": payload.AvatarId,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (repo *User) UpdateSettings(rctx *requests.RequestContext, userId uint, payload requests.SettingsRequest) error {
	result := repo.container.Config.ReadWriteDB.WithContext(rctx.Ctx).Model(&models.User{}).
		Where("id = ?", userId).
		Updates(map[string]interface{}{
			"currency_code":       payload.CurrencyCode,
			"number_format":       payload.NumberFormat,
			"decimal_places":      payload.DecimalPlaces,
			"email_notifications": payload.EmailNotifications,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

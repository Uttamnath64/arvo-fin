package repository

import (
	"errors"
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

func (repo *User) GetUserByUsernameOrEmail(username string, email string, user *models.User) error {
	return repo.container.Config.ReadOnlyDB.Model(&models.User{}).
		Where("username = ? or email = ?", username, strings.ToLower(email)).First(user).Error
}

func (repo *User) UsernameExists(username string) (bool, error) {
	var count int64

	err := repo.container.Config.ReadOnlyDB.Model(&models.User{}).
		Where("username = ?", username).Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (repo *User) EmailExists(email string) (bool, error) {
	var count int64

	err := repo.container.Config.ReadOnlyDB.Model(&models.User{}).
		Where("email = ?", strings.ToLower(email)).Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (repo *User) CreateUser(user *models.User) (uint, error) {
	err := repo.container.Config.ReadWriteDB.Create(user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (repo *User) UpdatePasswordByEmail(email, newPassword string) error {
	result := repo.container.Config.ReadWriteDB.Model(&models.User{}).
		Where("email = ?", email).
		Update("password", newPassword)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("User not found!")
	}
	return nil
}

func (repo *User) GetUser(userId uint, user *models.User) error {
	if err := repo.container.Config.ReadOnlyDB.Where("id = ?", userId).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *User) Get(userId uint) (*responses.MeResponse, error) {
	var user models.User
	var avatar models.Avatar
	var response responses.MeResponse

	query := repo.container.Config.ReadOnlyDB.Table(user.GetName()+" u").
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

func (repo *User) GetSettings(userId uint) (*responses.SettingsResponse, error) {

	var user models.User
	var currency models.Currency
	var response responses.SettingsResponse

	err := repo.container.Config.ReadOnlyDB.Table(user.GetName()+" u").
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

func (repo *User) Update(userId uint, payload requests.MeRequest) error {
	result := repo.container.Config.ReadWriteDB.Model(&models.User{}).
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
		return errors.New("User not found!")
	}
	return nil
}

func (repo *User) UpdateSettings(userId uint, payload requests.SettingsRequest) error {
	result := repo.container.Config.ReadWriteDB.Model(&models.User{}).
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
		return errors.New("User not found!")
	}
	return nil
}

package repository

import (
	"errors"
	"strings"

	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/storage"
)

type UserRepository struct {
	container *storage.Container
}

func NewUserRepository(container *storage.Container) *UserRepository {
	return &UserRepository{
		container: container,
	}
}

func (repo *UserRepository) GetUser(username string, email string, user *models.User) error {
	return repo.container.Config.ReadOnlyDB.Model(&models.User{}).
		Where("username = ? or email = ?", username, strings.ToLower(email)).First(user).Error
}

func (repo *UserRepository) UsernameExists(username string) (bool, error) {
	var count int64

	err := repo.container.Config.ReadOnlyDB.Model(&models.User{}).
		Where("username = ? or email = ?", username).Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (repo *UserRepository) EmailExists(email string) (bool, error) {
	var count int64

	err := repo.container.Config.ReadOnlyDB.Model(&models.User{}).
		Where("email = ?", strings.ToLower(email)).Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (repo *UserRepository) CreateUser(user *models.User) (uint, error) {
	err := repo.container.Config.ReadWriteDB.Create(user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil

}

func (repo *UserRepository) UpdatePasswordByEmail(email, newPassword string) error {
	result := repo.container.Config.ReadWriteDB.Model(&models.User{}).
		Where("email = ?", email).
		Update("password", newPassword)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("No user found!")
	}

	return nil
}

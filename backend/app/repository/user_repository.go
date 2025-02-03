package repository

import (
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

func (repo *UserRepository) GetUser(usernameOrEmail string, user *models.User) error {
	return repo.container.Config.ReadOnlyDB.Model(&models.User{}).
		Where("username = ? or email = ?", usernameOrEmail, strings.ToLower(usernameOrEmail)).First(user).Error
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

package repository

import (
	"time"

	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"gorm.io/gorm"
)

type TestUser struct {
	container *storage.Container
}

func NewTestUser(container *storage.Container) *TestUser {
	return &TestUser{
		container: container,
	}
}

func (repo *TestUser) GetUserByUsernameOrEmail(username string, email string, user *models.User) error {
	if username == "uttam.nath" || email == "uttam@example.com" {
		now := time.Now()
		*user = models.User{
			Model: gorm.Model{
				ID:        1,
				CreatedAt: now.AddDate(0, -1, 0),
				UpdatedAt: now,
			},
			Name:     "Uttam Nath",
			Email:    "uttam@example.com",
			Username: "uttam.nath",
			Password: "$2a$10$N7RKD8VqYHY4kbGWmfElBOs/wPfdnGldKAoRGOPa7ERbxEzeEOl1u",
		}
		return nil
	}
	return gorm.ErrRecordNotFound
}

func (repo *TestUser) UsernameExists(username string) (bool, error) {
	if username == "uttam.nath" {
		return true, nil
	}
	return false, nil
}

func (repo *TestUser) EmailExists(email string) (bool, error) {
	if email == "uttam@example.com" {
		return true, nil
	}
	return false, nil
}

func (repo *TestUser) CreateUser(user *models.User) (uint, error) {
	return 1, nil
}

func (repo *TestUser) UpdatePasswordByEmail(email, newPassword string) error {
	return nil
}

func (repo *TestUser) GetUser(userId uint, user *models.User) error {
	now := time.Now()
	*user = models.User{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: now.AddDate(0, -1, 0),
			UpdatedAt: now,
		},
		Name:     "Uttam Nath",
		Email:    "uttam@example.com",
		Username: "uttam.nath",
		Password: "$2a$10$N7RKD8VqYHY4kbGWmfElBOs/wPfdnGldKAoRGOPa7ERbxEzeEOl1u",
	}
	return nil
}

package repository

import (
	"time"

	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/responses"
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

func (repo *TestUser) GetUserByUsernameOrEmail(rctx *requests.RequestContext, username string, email string, user *models.User) error {
	if username == "uttam.nath" || email == "uttam@example.com" {
		now := time.Now()
		*user = models.User{
			BaseModel: models.BaseModel{
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

func (repo *TestUser) UsernameExists(rctx *requests.RequestContext, username string) error {
	if username == "uttam.nath" {
		return nil
	}
	return gorm.ErrRecordNotFound
}

func (repo *TestUser) EmailExists(rctx *requests.RequestContext, email string) error {
	if email == "uttam@example.com" {
		return nil
	}
	return gorm.ErrRecordNotFound
}

func (repo *TestUser) CreateUser(rctx *requests.RequestContext, user *models.User) (uint, error) {
	return 1, nil
}

func (repo *TestUser) UpdatePasswordByEmail(rctx *requests.RequestContext, email, newPassword string) error {
	return nil
}

func (repo *TestUser) GetUser(rctx *requests.RequestContext, userId uint, user *models.User) error {
	now := time.Now()
	*user = models.User{
		BaseModel: models.BaseModel{
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

func (repo *TestUser) Get(rctx *requests.RequestContext, userId uint) (*responses.MeResponse, error) {
	if userId != 1 {
		return nil, gorm.ErrRecordNotFound
	}
	return &responses.MeResponse{
		Id:         1,
		Name:       "uttam.nath",
		AvatarID:   8,
		AvatarIcon: "🧑",
	}, nil
}

func (repo *TestUser) GetSettings(rctx *requests.RequestContext, userId uint) (*responses.SettingsResponse, error) {
	if userId != 1 {
		return nil, gorm.ErrRecordNotFound
	}
	return &responses.SettingsResponse{
		Id:                 1,
		DecimalPlaces:      3,
		NumberFormat:       1,
		EmailNotifications: true,
		CurrencyCode:       "INR",
		CurrencyName:       "India",
		CurrencySymbol:     "₹",
	}, nil
}

func (repo *TestUser) Update(rctx *requests.RequestContext, userId uint, payload requests.MeRequest) error {
	if userId != 1 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (repo *TestUser) UpdateSettings(rctx *requests.RequestContext, userId uint, payload requests.SettingsRequest) error {
	if userId != 1 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

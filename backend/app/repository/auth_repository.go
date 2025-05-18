package repository

import (
	"errors"

	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/storage"
)

type Auth struct {
	container *storage.Container
}

func NewAuth(container *storage.Container) *Auth {
	return &Auth{
		container: container,
	}
}

func (repo *Auth) GetSessionByUser(userId uint, userType commonType.UserType, signedToken string) (*models.Session, error) {
	var session models.Session
	err := repo.container.Config.ReadOnlyDB.Where("user_id = ? AND user_type = ? AND token = ?", userId, userType, signedToken).First(&session).Error
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (repo *Auth) GetSessionByRefreshToken(refreshToken string, userType commonType.UserType) (*models.Session, error) {
	var session models.Session
	err := repo.container.Config.ReadOnlyDB.Where("refresh_token = ? AND user_type = ?", refreshToken, userType).First(&session).Error
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (repo *Auth) CreateSession(session *models.Session) error {
	return repo.container.Config.ReadOnlyDB.Create(session).Error
}

func (repo *Auth) UpdateSession(id uint, refreshToken string, expiresAt int64) error {
	result := repo.container.Config.ReadWriteDB.Model(&models.Portfolio{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"refresh_token": refreshToken,
			"expires_at":    expiresAt,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("Session not found!")
	}
	return nil
}

package repository

import (
	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"gorm.io/gorm"
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

func (repo *Auth) CreateSession(session *models.Session) (uint, error) {
	err := repo.container.Config.ReadWriteDB.Create(session).Error
	if err != nil {
		return 0, err
	}
	return session.ID, nil
}

func (repo *Auth) DeleteSession(sessionID uint) error {
	return repo.container.Config.ReadWriteDB.Unscoped().Where("id = ?", sessionID).Delete(&models.Session{}).Error
}

func (repo *Auth) UpdateSession(id uint, refreshToken string, expiresAt int64) error {
	result := repo.container.Config.ReadWriteDB.Model(&models.Session{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"refresh_token": refreshToken,
			"expires_at":    expiresAt,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

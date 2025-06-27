package repository

import (
	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/requests"
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

func (repo *Auth) GetSessionByUser(rctx *requests.RequestContext, userId uint, userType commonType.UserType, signedToken string) (*models.Session, error) {
	var session models.Session
	err := repo.container.Config.ReadOnlyDB.WithContext(rctx.Ctx).Where("user_id = ? AND user_type = ? AND token = ?", userId, userType, signedToken).First(&session).Error
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (repo *Auth) GetSessionByRefreshToken(rctx *requests.RequestContext, refreshToken string, userType commonType.UserType) (*models.Session, error) {
	var session models.Session
	err := repo.container.Config.ReadOnlyDB.WithContext(rctx.Ctx).Where("refresh_token = ? AND user_type = ?", refreshToken, userType).First(&session).Error
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (repo *Auth) CreateSession(rctx *requests.RequestContext, session *models.Session) (uint, error) {
	err := repo.container.Config.ReadWriteDB.WithContext(rctx.Ctx).Create(session).Error
	if err != nil {
		return 0, err
	}
	return session.ID, nil
}

func (repo *Auth) DeleteSession(rctx *requests.RequestContext, sessionID uint) error {
	return repo.container.Config.ReadWriteDB.WithContext(rctx.Ctx).Unscoped().Where("id = ?", sessionID).Delete(&models.Session{}).Error
}

func (repo *Auth) UpdateSession(rctx *requests.RequestContext, id uint, refreshToken string, expiresAt int64) error {
	result := repo.container.Config.ReadWriteDB.WithContext(rctx.Ctx).Model(&models.Session{}).
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

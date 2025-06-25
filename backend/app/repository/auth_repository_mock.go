package repository

import (
	"time"

	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"gorm.io/gorm"
)

type TestAuth struct {
	container *storage.Container
}

func NewTestAuth(container *storage.Container) *TestAuth {
	return &TestAuth{
		container: container,
	}
}

func (repo *TestAuth) GetSessionByUser(userId uint, userType commonType.UserType, signedToken string) (*models.Session, error) {
	if userId == 1 && userType == commonType.UserTypeUser {
		session := models.Session{
			BaseModel: models.BaseModel{
				ID:        1,
				CreatedAt: time.Now().Add(-1 * time.Hour),
				UpdatedAt: time.Now(),
			},
			UserID:       101,
			UserType:     commonType.UserTypeUser,
			DeviceInfo:   "Chrome on Windows 10",
			IPAddress:    "192.168.1.100",
			RefreshToken: "mock-refresh-token-abc123",
			ExpiresAt:    time.Now().Add(7 * 24 * time.Hour).Unix(),
		}
		return &session, nil
	}
	return nil, gorm.ErrRecordNotFound
}

func (repo *TestAuth) GetSessionByRefreshToken(refreshToken string, userType commonType.UserType) (*models.Session, error) {
	if refreshToken != "" && userType == commonType.UserTypeUser {
		session := models.Session{
			BaseModel: models.BaseModel{
				ID:        1,
				CreatedAt: time.Now().Add(-1 * time.Hour),
				UpdatedAt: time.Now(),
			},
			UserID:       101,
			UserType:     commonType.UserTypeUser,
			DeviceInfo:   "Chrome on Windows 10",
			IPAddress:    "192.168.1.100",
			RefreshToken: "mock-refresh-token-abc123",
			ExpiresAt:    7382945133,
		}
		return &session, nil
	}
	return nil, gorm.ErrRecordNotFound
}

func (repo *TestAuth) CreateSession(session *models.Session) (uint, error) {
	return 1, nil
}

func (repo *TestAuth) DeleteSession(sessionID uint) error {
	return nil
}

func (repo *TestAuth) UpdateSession(id uint, refreshToken string, expiresAt int64) error {
	return nil
}

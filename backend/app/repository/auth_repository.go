package repository

import (
	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"gorm.io/gorm"
)

type AuthRepository struct {
	container *storage.Container
}

func NewAuthRepository(container *storage.Container) *AuthRepository {
	return &AuthRepository{
		container: container,
	}
}

func (repo *AuthRepository) GetTokenByReference(referenceID uint, userType commonType.UserType, signedToken string) (*models.Token, error) {
	var token models.Token
	err := repo.container.Config.ReadOnlyDB.Where("reference_id = ? AND user_type = ? AND token = ?", referenceID, userType, signedToken).First(&token).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // No token found
		}
		return nil, err // Other errors
	}
	return &token, nil
}

func (repo *AuthRepository) GetTokenByRefreshToken(refreshToken uint, userType commonType.UserType) (*models.Token, error) {
	var token models.Token
	err := repo.container.Config.ReadOnlyDB.Where("refresh_token = ? AND user_type = ?", refreshToken, userType).First(&token).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // No token found
		}
		return nil, err // Other errors
	}
	return &token, nil
}

func (repo *AuthRepository) CreateToken(token *models.Token) error {
	return repo.container.Config.ReadOnlyDB.Create(token).Error
}

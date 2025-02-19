package repository

import (
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

func (repo *Auth) GetTokenByReference(referenceID uint, userType commonType.UserType, signedToken string) (*models.Token, error) {
	var token models.Token
	err := repo.container.Config.ReadOnlyDB.Where("reference_id = ? AND user_type = ? AND token = ?", referenceID, userType, signedToken).First(&token).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (repo *Auth) GetTokenByRefreshToken(refreshToken uint, userType commonType.UserType) (*models.Token, error) {
	var token models.Token
	err := repo.container.Config.ReadOnlyDB.Where("refresh_token = ? AND user_type = ?", refreshToken, userType).First(&token).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (repo *Auth) CreateToken(token *models.Token) error {
	return repo.container.Config.ReadOnlyDB.Create(token).Error
}

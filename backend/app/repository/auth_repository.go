package repository

import (
	"github.com/Uttamnath64/arvo-fin/app/config"
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/pkg/logger"
	"gorm.io/gorm"
)

type AuthRepository struct {
	config *config.Config
	logger *logger.Logger
}

func NewAuthRepository(config *config.Config, logger *logger.Logger) *AuthRepository {
	return &AuthRepository{
		config: config,
		logger: logger,
	}
}

func (repo *AuthRepository) GetTokenByReference(referenceID uint, userType byte, signedToken string) (*models.Token, error) {
	var token models.Token
	err := repo.config.ReadOnlyDB.Where("referenceId = ? AND userType = ? AND token = ?", referenceID, userType, signedToken).First(&token).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // No token found
		}
		return nil, err // Other errors
	}
	return &token, nil
}

func (repo *AuthRepository) AddToken(token *models.Token) error {
	return repo.config.ReadOnlyDB.Create(token).Error
}

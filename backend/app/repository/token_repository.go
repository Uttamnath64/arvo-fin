package repository

import (
	"github.com/Uttamnath64/arvo-fin/app/config"
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/pkg/logger"
	"gorm.io/gorm"
)

type TokenRepository struct {
	config *config.Config
	logger *logger.Logger
}

func NewTokenRepository(config *config.Config, logger *logger.Logger) *TokenRepository {
	return &TokenRepository{
		config: config,
		logger: logger,
	}
}

func (r *TokenRepository) GetTokenByReference(referenceID uint, userType byte, signedToken string) (*models.Token, error) {
	var token models.Token
	err := r.config.ReadOnlyDB.Where("referenceId = ? AND userType = ? AND token = ?", referenceID, userType, signedToken).First(&token).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // No token found
		}
		return nil, err // Other errors
	}
	return &token, nil
}

func (r *TokenRepository) AddToken(token *models.Token) error {
	return r.config.ReadOnlyDB.Create(token).Error
}

package services

import (
	"errors"
	"time"

	"github.com/Uttamnath64/arvo-fin/app/config"
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/repository"
	"github.com/Uttamnath64/arvo-fin/pkg/logger"
)

type AuthService struct {
	config    *config.Config
	logger    *logger.Logger
	tokenRepo *repository.TokenRepository
}

func NewAuthService(config *config.Config, logger *logger.Logger) *AuthService {
	return &AuthService{
		config:    config,
		logger:    logger,
		tokenRepo: repository.NewTokenRepository(config, logger),
	}
}

func (t *AuthService) IsValidRefreshToken(referenceID uint, userType byte, signedToken string) error {
	token, err := t.tokenRepo.GetTokenByReference(referenceID, userType, signedToken)
	if err != nil {
		return err
	}

	// Check if token exists
	if token == nil {
		return errors.New("Refresh token not found!")
	}

	if token.ExpiresAt < time.Now().Unix() {
		return errors.New("Refresh token is expired!")
	}

	return nil
}

func (t *AuthService) AddToken(token *models.Token) error {
	return t.tokenRepo.AddToken(token)
}

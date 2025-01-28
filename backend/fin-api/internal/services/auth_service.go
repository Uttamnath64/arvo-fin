package services

import (
	"github.com/Uttamnath64/arvo-fin/app/config"
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/services"
	"github.com/Uttamnath64/arvo-fin/pkg/logger"
)

type AuthService struct {
	config      *config.Config
	logger      *logger.Logger
	authService *services.AuthService
}

func NewAuthService(config *config.Config, logger *logger.Logger) *AuthService {
	return &AuthService{
		config:      config,
		logger:      logger,
		authService: services.NewAuthService(config, logger),
	}
}

func (service *AuthService) IsValidRefreshToken(referenceID uint, userType byte, signedToken string) error {
	return service.authService.IsValidRefreshToken(referenceID, userType, signedToken)
}

func (service *AuthService) AddToken(token *models.Token) error {
	return service.authService.AddToken(token)
}

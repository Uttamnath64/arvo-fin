package services

import (
	"github.com/Uttamnath64/arvo-fin/app/config"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/responses"
	"github.com/Uttamnath64/arvo-fin/app/services"
	"github.com/Uttamnath64/arvo-fin/pkg/logger"
)

type AuthService struct {
	config      *config.Config
	logger      *logger.Logger
	authService *services.AuthService
	env         *config.AppEnv
}

func NewAuthService(config *config.Config, logger *logger.Logger, env *config.AppEnv) *AuthService {
	return &AuthService{
		config:      config,
		logger:      logger,
		authService: services.NewAuthService(config, logger, env),
		env:         env,
	}
}

func (service *AuthService) Login(ip string, payload requests.LoginRequest) *responses.ServiceResponse {
	return service.authService.Login(ip, payload)
}

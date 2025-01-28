package services

import (
	"github.com/Uttamnath64/arvo-fin/app/config"
	"github.com/Uttamnath64/arvo-fin/app/repository"
	"github.com/Uttamnath64/arvo-fin/pkg/logger"
)

type UserService struct {
	config *config.Config
	logger *logger.Logger
	repo   *repository.UserRepository
}

func NewUserService(config *config.Config, logger *logger.Logger) *UserService {
	return &UserService{
		config: config,
		logger: logger,
		repo:   repository.NewUserRepository(config, logger),
	}
}

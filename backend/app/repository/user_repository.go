package repository

import (
	"github.com/Uttamnath64/arvo-fin/app/config"
	"github.com/Uttamnath64/arvo-fin/pkg/logger"
)

type UserRepository struct {
	config *config.Config
	logger *logger.Logger
}

func NewUserRepository(config *config.Config, logger *logger.Logger) *UserRepository {
	return &UserRepository{
		config: config,
		logger: logger,
	}
}

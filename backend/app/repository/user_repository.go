package repository

import (
	"strings"

	"github.com/Uttamnath64/arvo-fin/app/config"
	"github.com/Uttamnath64/arvo-fin/app/models"
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

func (repo *UserRepository) GetUser(usernameOrEmail string, user *models.User) error {
	return repo.config.ReadOnlyDB.Model(&models.User{}).
		Where("username = ? or email = ?", usernameOrEmail, strings.ToLower(usernameOrEmail)).First(user).Error
}

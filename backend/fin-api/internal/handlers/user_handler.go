package handlers

import (
	"github.com/Uttamnath64/arvo-fin/app/config"
	"github.com/Uttamnath64/arvo-fin/fin-api/internal/services"
	"github.com/Uttamnath64/arvo-fin/pkg/logger"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	config      *config.Config
	logger      *logger.Logger
	userService *services.UserService
}

func NewUserHandler(config *config.Config, logger *logger.Logger) *UserHandler {
	return &UserHandler{
		config:      config,
		logger:      logger,
		userService: services.NewUserService(config, logger),
	}
}

func (handler *UserHandler) GetUsersHandler(c *gin.Context) {

}

func (handler *UserHandler) CreateUserHandler(c *gin.Context) {

}

func (handler *UserHandler) GetUserByIDHandler(c *gin.Context) {

}

func (handler *UserHandler) UpdateUserHandler(c *gin.Context) {

}

func (handler *UserHandler) DeleteUserHandler(c *gin.Context) {

}

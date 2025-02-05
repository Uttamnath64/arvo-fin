package handlers

import (
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"github.com/Uttamnath64/arvo-fin/fin-api/internal/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	container   *storage.Container
	userService *services.UserService
}

func NewUserHandler(container *storage.Container) *UserHandler {
	return &UserHandler{
		container:   container,
		userService: services.NewUserService(container),
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

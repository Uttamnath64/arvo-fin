package services

import (
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/responses"
	"github.com/Uttamnath64/arvo-fin/app/services"
	"github.com/Uttamnath64/arvo-fin/app/storage"
)

type User struct {
	container   *storage.Container
	userService services.UserService
}

func NewUser(container *storage.Container) *User {
	return &User{
		container:   container,
		userService: services.NewUser(container),
	}
}

func (service *User) Get(rctx *requests.RequestContext, userId uint) responses.ServiceResponse {
	return service.userService.Get(rctx, userId)
}

func (service *User) GetSettings(rctx *requests.RequestContext, userId uint) responses.ServiceResponse {
	return service.userService.GetSettings(rctx, userId)
}

func (service *User) Update(rctx *requests.RequestContext, payload requests.MeRequest, userId uint) responses.ServiceResponse {
	return service.userService.Update(rctx, payload, userId)
}

func (service *User) UpdateSettings(rctx *requests.RequestContext, payload requests.SettingsRequest, userId uint) responses.ServiceResponse {
	return service.userService.UpdateSettings(rctx, payload, userId)
}

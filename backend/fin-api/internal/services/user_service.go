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

func (service *User) Get(userId uint) responses.ServiceResponse {
	return service.userService.Get(userId)
}

func (service *User) GetSettings(userId uint) responses.ServiceResponse {
	return service.userService.GetSettings(userId)
}

func (service *User) Update(payload requests.MeRequest, userId uint) responses.ServiceResponse {
	return service.userService.Update(payload, userId)
}

func (service *User) UpdateSettings(payload requests.SettingsRequest, userId uint) responses.ServiceResponse {
	return service.userService.UpdateSettings(payload, userId)
}

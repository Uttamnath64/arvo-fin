package services

import (
	"github.com/Uttamnath64/arvo-fin/app/repository"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/responses"
	"github.com/Uttamnath64/arvo-fin/app/storage"
)

type User struct {
	container *storage.Container
	userRepo  repository.UserRepository
}

func NewUser(container *storage.Container) *User {
	return &User{
		container: container,
		userRepo:  repository.NewUser(container),
	}
}

func (service *User) Get(userId uint) responses.ServiceResponse {
	return responses.ServiceResponse{}
}

func (service *User) GetSettings(userId uint) responses.ServiceResponse {
	return responses.ServiceResponse{}
}

func (service *User) Update(payload requests.SettingsRequest, userId uint) responses.ServiceResponse {
	return responses.ServiceResponse{}
}

func (service *User) UpdateSettings(payload requests.SettingsRequest, userId uint) responses.ServiceResponse {
	return responses.ServiceResponse{}
}

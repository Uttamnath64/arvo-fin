package services

import (
	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/responses"
	"github.com/Uttamnath64/arvo-fin/app/services"
	"github.com/Uttamnath64/arvo-fin/app/storage"
)

type Avatar struct {
	container     *storage.Container
	avatarService services.AvatarService
}

func NewAvatar(container *storage.Container) *Avatar {
	return &Avatar{
		container:     container,
		avatarService: services.NewAvatar(container),
	}
}

func (service *Avatar) Get(id uint) responses.ServiceResponse {
	return service.avatarService.Get(id)
}

func (service *Avatar) GetAvatarsByType(avatarType commonType.AvatarType) responses.ServiceResponse {
	return service.avatarService.GetAvatarsByType(avatarType)
}

func (service *Avatar) Create(payload requests.AvatarRequest) responses.ServiceResponse {
	return service.avatarService.Creatre(payload)
}

func (service *Avatar) Update(id uint, payload requests.AvatarRequest) responses.ServiceResponse {
	return service.avatarService.Update(id, payload)
}

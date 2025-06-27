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

func (service *Avatar) Get(rctx *requests.RequestContext, id uint) responses.ServiceResponse {
	return service.avatarService.Get(rctx, id)
}

func (service *Avatar) GetAvatarsByType(rctx *requests.RequestContext, avatarType commonType.AvatarType) responses.ServiceResponse {
	return service.avatarService.GetAvatarsByType(rctx, avatarType)
}

func (service *Avatar) Create(rctx *requests.RequestContext, payload requests.AvatarRequest) responses.ServiceResponse {
	return service.avatarService.Creatre(rctx, payload)
}

func (service *Avatar) Update(rctx *requests.RequestContext, id uint, payload requests.AvatarRequest) responses.ServiceResponse {
	return service.avatarService.Update(rctx, id, payload)
}

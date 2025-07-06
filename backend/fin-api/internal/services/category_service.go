package services

import (
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/responses"
	"github.com/Uttamnath64/arvo-fin/app/services"
	"github.com/Uttamnath64/arvo-fin/app/storage"
)

type Category struct {
	container *storage.Container
	service   services.CategoryService
}

func NewCategory(container *storage.Container) *Category {
	return &Category{
		container: container,
		service:   services.NewCategory(container),
	}
}

func (service *Category) Get(rctx *requests.RequestContext, id uint) responses.ServiceResponse {
	return service.service.Get(rctx, id)
}
func (service *Category) GetList(rctx *requests.RequestContext, portfolioId, userId uint) responses.ServiceResponse {
	return service.service.GetList(rctx, portfolioId, userId)
}

func (service *Category) Create(rctx *requests.RequestContext, payload requests.CategoryRequest) responses.ServiceResponse {
	return service.service.Create(rctx, payload)
}

func (service *Category) Update(rctx *requests.RequestContext, id uint, payload requests.CategoryRequest) responses.ServiceResponse {
	return service.service.Update(rctx, id, payload)
}

func (service *Category) Delete(rctx *requests.RequestContext, portfolioId, id uint) responses.ServiceResponse {
	return service.service.Delete(rctx, portfolioId, id)
}

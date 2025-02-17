package services

import (
	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/responses"
	commonServices "github.com/Uttamnath64/arvo-fin/app/services"
	"github.com/Uttamnath64/arvo-fin/app/storage"
)

type Portfolio struct {
	container        *storage.Container
	portfolioService *commonServices.Portfolio
}

func NewPortfolio(container *storage.Container) *Portfolio {
	return &Portfolio{
		container:        container,
		portfolioService: commonServices.NewPortfolio(container),
	}
}

func (service *Portfolio) GetList(userId uint, userType commonType.UserType) responses.ServiceResponse {
	return service.portfolioService.GetList(userId, userType)
}

func (service *Portfolio) Get(id, userId uint, userType commonType.UserType) responses.ServiceResponse {
	return service.portfolioService.Get(id, userId, userType)
}

func (service *Portfolio) Add(payload requests.PortfolioRequest, userId uint) responses.ServiceResponse {
	return service.portfolioService.Add(payload, userId)
}

func (service *Portfolio) Update(id, userId uint, payload requests.PortfolioRequest) responses.ServiceResponse {
	return service.portfolioService.Update(id, userId, payload)
}

func (service *Portfolio) Delete(id, userId uint) responses.ServiceResponse {
	return service.portfolioService.Delete(id, userId)
}

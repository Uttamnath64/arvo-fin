package services

import (
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/responses"
	"github.com/Uttamnath64/arvo-fin/app/services"
	"github.com/Uttamnath64/arvo-fin/app/storage"
)

type Account struct {
	container      *storage.Container
	accountService services.AccountService
}

func NewAccount(container *storage.Container) *Account {
	return &Account{
		container:      container,
		accountService: services.NewAccount(container),
	}
}

func (service *Account) Get(rctx *requests.RequestContext, id uint) responses.ServiceResponse {
	return service.accountService.Get(rctx, id)
}
func (service *Account) GetList(rctx *requests.RequestContext, portfolioId, userId uint) responses.ServiceResponse {
	return service.accountService.GetList(rctx, portfolioId, userId)
}

func (service *Account) AccountTypes(rctx *requests.RequestContext) responses.ServiceResponse {
	return service.accountService.AccountTypes(rctx)
}

func (service *Account) Create(rctx *requests.RequestContext, userId uint, payload requests.AccountRequest) responses.ServiceResponse {
	return service.accountService.Create(rctx, userId, payload)
}

func (service *Account) Update(rctx *requests.RequestContext, id, userId uint, payload requests.AccountUpdateRequest) responses.ServiceResponse {
	return service.accountService.Update(rctx, id, userId, payload)
}

func (service *Account) Delete(rctx *requests.RequestContext, id, userId uint) responses.ServiceResponse {
	return service.accountService.Delete(rctx, id, userId)
}

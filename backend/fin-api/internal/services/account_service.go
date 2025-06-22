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

func (service *Account) Get(id uint) responses.ServiceResponse {
	return service.accountService.Get(id)
}
func (service *Account) GetList(portfolioId, userId uint) responses.ServiceResponse {
	return service.accountService.GetList(portfolioId, userId)
}

func (service *Account) AccountTypes() responses.ServiceResponse {
	return service.accountService.AccountTypes()
}

func (service *Account) Create(userId uint, payload requests.AccountRequest) responses.ServiceResponse {
	return service.accountService.Create(userId, payload)
}

func (service *Account) Update(id, userId uint, payload requests.AccountUpdateRequest) responses.ServiceResponse {
	return service.accountService.Update(id, userId, payload)
}

func (service *Account) Delete(id, userId uint) responses.ServiceResponse {
	return service.accountService.Delete(id, userId)
}

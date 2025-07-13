package services

import (
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/responses"
	"github.com/Uttamnath64/arvo-fin/app/services"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"github.com/Uttamnath64/arvo-fin/pkg/pagination"
)

type Transaction struct {
	container *storage.Container
	service   services.TransactionService
}

func NewTransaction(container *storage.Container) *Transaction {
	return &Transaction{
		container: container,
		service:   services.NewTransaction(container),
	}
}

func (service *Transaction) Get(rctx *requests.RequestContext, id uint) responses.ServiceResponse {
	return service.service.Get(rctx, id)
}
func (service *Transaction) GetList(rctx *requests.RequestContext, transactionQuery requests.TransactionQuery, pagination pagination.Pagination) responses.ServiceResponse {
	return service.service.GetList(rctx, transactionQuery, pagination)
}

func (service *Transaction) Create(rctx *requests.RequestContext, payload requests.TransactionRequest) responses.ServiceResponse {
	return service.service.Create(rctx, payload)
}

func (service *Transaction) Update(rctx *requests.RequestContext, id uint, payload requests.TransactionRequest) responses.ServiceResponse {
	return service.service.Update(rctx, id, payload)
}

func (service *Transaction) Delete(rctx *requests.RequestContext, id uint) responses.ServiceResponse {
	return service.service.Delete(rctx, id)
}

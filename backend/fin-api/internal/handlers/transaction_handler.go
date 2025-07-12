package handlers

import (
	"net/http"
	"strconv"

	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/responses"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"github.com/Uttamnath64/arvo-fin/fin-api/internal/services"
	"github.com/Uttamnath64/arvo-fin/pkg/pagination"
	"github.com/Uttamnath64/arvo-fin/pkg/query"
	"github.com/gin-gonic/gin"
)

type Transaction struct {
	container *storage.Container
	service   *services.Transaction
}

func NewTransaction(container *storage.Container) *Transaction {
	return &Transaction{
		container: container,
		service:   services.NewTransaction(container),
	}
}

func (handler *Transaction) Get(c *gin.Context) {

	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: "Invalid transaction id!",
		})
		return
	}

	serviceResponse := handler.service.Get(rctx, uint(id))
	if isErrorResponse(c, serviceResponse) {
		return
	}

	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: serviceResponse.Data,
	})
}

func (handler *Transaction) GetList(c *gin.Context) {
	var isRequired bool
	var transactionQuery requests.TransactionQuery

	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}

	transactionQuery.UserId = rctx.UserID
	isRequired = true

	if rctx.UserType == commonType.UserTypeAdmin {
		isRequired = false
		transactionQuery.UserId, ok = query.QId(c, "userId", false)
		if !ok {
			return
		}
	}

	// portfolioId
	transactionQuery.PortfolioId, ok = query.QId(c, "portfolioId", isRequired)
	if !ok {
		return
	}

	// accountId
	transactionQuery.AccountId, ok = query.QId(c, "accountId", false)
	if !ok {
		return
	}

	// categoryId
	transactionQuery.CategoryId, ok = query.QId(c, "categoryId", false)
	if !ok {
		return
	}

	// dateFrom and dateTo
	transactionQuery.DateFrom, transactionQuery.DateTo, ok = query.QDateTimeRange(c, "dateFrom", "dateTo", false)
	if !ok {
		return
	}

	// search
	transactionQuery.Search = c.Query("search")

	// type
	tType := c.Query("type")
	if tType != "" {
		typeInt, err := strconv.Atoi(tType)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.ApiResponse{
				Status:  false,
				Message: "Invalid type!",
			})
		}
		tType := commonType.TransactionType(typeInt)
		if !commonType.OrderType(transactionQuery.Order).IsValid() {
			c.JSON(http.StatusBadRequest, responses.ApiResponse{
				Status:  false,
				Message: "Invalid type!",
			})
			return
		}
		transactionQuery.Type = &tType
	}

	// order
	transactionQuery.Order = commonType.OrderType(c.DefaultQuery("order", "asc"))
	if !commonType.OrderType(transactionQuery.Order).IsValid() {
		c.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: "Invalid order type. Use 'asc' or 'desc'",
		})
		return
	}

	pagination := pagination.NewPagination(c)

	serviceResponse := handler.service.GetList(rctx, transactionQuery, pagination)
	if isErrorResponse(c, serviceResponse) {
		return
	}

	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: serviceResponse.Data,
	})
}

func (handler *Transaction) Create(c *gin.Context) {

	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}

	var payload requests.TransactionRequest
	if !bindAndValidateJson(c, &payload) {
		return
	}

	if rctx.UserType != commonType.UserTypeUser {
		c.JSON(http.StatusForbidden, responses.ApiResponse{
			Status:  false,
			Message: "Only user can update transaction!",
		})
		return
	}

	serviceResponse := handler.service.Create(rctx, payload)
	if isErrorResponse(c, serviceResponse) {
		return
	}

	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: serviceResponse.Data,
	})
}

func (handler *Transaction) Update(c *gin.Context) {

	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}

	var payload requests.TransactionRequest
	if !bindAndValidateJson(c, &payload) {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: "Invalid transaction id!",
		})
		return
	}

	if rctx.UserType != commonType.UserTypeUser {
		c.JSON(http.StatusForbidden, responses.ApiResponse{
			Status:  false,
			Message: "Only user can transaction account!",
		})
		return
	}

	serviceResponse := handler.service.Update(rctx, uint(id), payload)
	if isErrorResponse(c, serviceResponse) {
		return
	}

	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: serviceResponse.Data,
	})
}

func (handler *Transaction) Delete(c *gin.Context) {

	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: "Invalid transaction id!",
		})
		return
	}

	if rctx.UserType != commonType.UserTypeUser {
		c.JSON(http.StatusForbidden, responses.ApiResponse{
			Status:  false,
			Message: "Only user can update transaction!",
		})
		return
	}

	serviceResponse := handler.service.Delete(rctx, uint(id))
	if isErrorResponse(c, serviceResponse) {
		return
	}

	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: serviceResponse.Data,
	})
}

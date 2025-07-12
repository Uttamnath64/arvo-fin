package handlers

import (
	"net/http"
	"strconv"

	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/responses"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"github.com/Uttamnath64/arvo-fin/fin-api/internal/services"
	"github.com/gin-gonic/gin"
)

type Transaction struct {
	container *storage.Container
	service   *services.Account
}

func NewTransaction(container *storage.Container) *Transaction {
	return &Transaction{
		container: container,
		service:   services.NewAccount(container),
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
	var queryParams map[string]interface{}
	var isRequired bool
	var uid uint

	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}

	if rctx.UserType == commonType.UserTypeAdmin {
		uid, ok := getQueryParamId(c, "userId", false)
		if !ok {
			return
		}
		queryParams["userId"] = uid
	} else {
		queryParams["userId"] = rctx.UserID
		isRequired = true
	}

	uid, ok = getQueryParamId(c, "portfolioId", isRequired)
	if !ok {
		return
	}
	queryParams["portfolioId"] = uid

	uid, ok = getQueryParamId(c, "accountId", false)
	if !ok {
		return
	}
	queryParams["accountId"] = uid

	uid, ok = getQueryParamId(c, "categoryId", false)
	if !ok {
		return
	}
	queryParams["categoryId"] = uid

	serviceResponse := handler.service.GetList(rctx, queryParams)
	if isErrorResponse(c, serviceResponse) {
		return
	}

	search := c.Query("search")
	tType := c.Query("type")
	dateFrom := c.Query("dateFrom")
	dateTo := c.Query("dateTo")
	order := c.DefaultQuery("order", "desc")

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

	var payload requests.AccountRequest
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

	serviceResponse := handler.service.Create(rctx, rctx.UserID, payload)
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

	var payload requests.AccountUpdateRequest
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

	serviceResponse := handler.service.Update(rctx, uint(id), rctx.UserID, payload)
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

	serviceResponse := handler.service.Delete(rctx, uint(id), rctx.UserID)
	if isErrorResponse(c, serviceResponse) {
		return
	}

	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: serviceResponse.Data,
	})
}

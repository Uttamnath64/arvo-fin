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

type Account struct {
	container      *storage.Container
	accountService *services.Account
}

func NewAccount(container *storage.Container) *Account {
	return &Account{
		container:      container,
		accountService: services.NewAccount(container),
	}
}

func (handler *Account) Get(c *gin.Context) {

	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: "Invalid account id!",
		})
		return
	}

	serviceResponse := handler.accountService.Get(rctx, uint(id))
	if isErrorResponse(c, serviceResponse) {
		return
	}

	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: serviceResponse.Data,
	})
}

func (handler *Account) GetList(c *gin.Context) {

	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}
	userId := rctx.UserID

	portfolioId, err := strconv.Atoi(c.Param("portfolioId"))
	if err != nil || portfolioId <= 0 {
		c.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: "Invalid portfolio id!",
		})
		return
	}

	// User is admin then
	if rctx.UserType == commonType.UserTypeAdmin {
		id, err := strconv.Atoi(c.Param("userId"))
		if err != nil || id <= 0 {
			c.JSON(http.StatusBadRequest, responses.ApiResponse{
				Status:  false,
				Message: "Invalid user id!",
			})
			return
		}
		userId = uint(id)
	}

	serviceResponse := handler.accountService.GetList(rctx, uint(portfolioId), userId)
	if isErrorResponse(c, serviceResponse) {
		return
	}

	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: serviceResponse.Data,
	})
}

func (handler *Account) AccountTypes(c *gin.Context) {

	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}

	serviceResponse := handler.accountService.AccountTypes(rctx)
	if isErrorResponse(c, serviceResponse) {
		return
	}

	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: serviceResponse.Data,
	})
}

func (handler *Account) Create(c *gin.Context) {

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
			Message: "Only user can update account!",
		})
		return
	}

	serviceResponse := handler.accountService.Create(rctx, rctx.UserID, payload)
	if isErrorResponse(c, serviceResponse) {
		return
	}

	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: serviceResponse.Data,
	})
}

func (handler *Account) Update(c *gin.Context) {

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
			Message: "Invalid account id!",
		})
		return
	}

	if rctx.UserType != commonType.UserTypeUser {
		c.JSON(http.StatusForbidden, responses.ApiResponse{
			Status:  false,
			Message: "Only user can update account!",
		})
		return
	}

	serviceResponse := handler.accountService.Update(rctx, uint(id), rctx.UserID, payload)
	if isErrorResponse(c, serviceResponse) {
		return
	}

	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: serviceResponse.Data,
	})
}

func (handler *Account) Delete(c *gin.Context) {

	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: "Invalid account id!",
		})
		return
	}

	if rctx.UserType != commonType.UserTypeUser {
		c.JSON(http.StatusForbidden, responses.ApiResponse{
			Status:  false,
			Message: "Only user can update account!",
		})
		return
	}

	serviceResponse := handler.accountService.Delete(rctx, uint(id), rctx.UserID)
	if isErrorResponse(c, serviceResponse) {
		return
	}

	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: serviceResponse.Data,
	})
}

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

func (handler *Account) Get(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: "Invalid account id!",
		})
		return
	}

	serviceResponse := handler.accountService.Get(uint(id))
	if isErrorResponse(ctx, serviceResponse) {
		return
	}

	ctx.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: serviceResponse.Data,
	})
}

func (handler *Account) GetList(ctx *gin.Context) {

	userInfo, ok := getUserInfo(ctx)
	if !ok {
		return
	}
	userId := userInfo.userId

	portfolioId, err := strconv.Atoi(ctx.Param("portfolioId"))
	if err != nil || portfolioId <= 0 {
		ctx.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: "Invalid portfolio id!",
		})
		return
	}

	// User is admin then
	if userInfo.userType == commonType.UserTypeAdmin {
		id, err := strconv.Atoi(ctx.Param("userId"))
		if err != nil || id <= 0 {
			ctx.JSON(http.StatusBadRequest, responses.ApiResponse{
				Status:  false,
				Message: "Invalid user id!",
			})
			return
		}
		userId = uint(id)
	}

	serviceResponse := handler.accountService.GetList(uint(portfolioId), userId)
	if isErrorResponse(ctx, serviceResponse) {
		return
	}

	ctx.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: serviceResponse.Data,
	})
}

func (handler *Account) AccountTypes(ctx *gin.Context) {

	serviceResponse := handler.accountService.AccountTypes()
	if isErrorResponse(ctx, serviceResponse) {
		return
	}

	ctx.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: serviceResponse.Data,
	})
}

func (handler *Account) Create(ctx *gin.Context) {

	var payload requests.AccountRequest
	if !bindAndValidateJson(ctx, &payload) {
		return
	}

	userInfo, ok := getUserInfo(ctx)
	if !ok {
		return
	}

	if userInfo.userType != commonType.UserTypeUser {
		ctx.JSON(http.StatusForbidden, responses.ApiResponse{
			Status:  false,
			Message: "Only user can update account!",
		})
		return
	}

	serviceResponse := handler.accountService.Create(userInfo.userId, payload)
	if isErrorResponse(ctx, serviceResponse) {
		return
	}

	ctx.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: serviceResponse.Data,
	})
}

func (handler *Account) Update(ctx *gin.Context) {

	var payload requests.AccountUpdateRequest
	if !bindAndValidateJson(ctx, &payload) {
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: "Invalid account id!",
		})
		return
	}

	userInfo, ok := getUserInfo(ctx)
	if !ok {
		return
	}

	if userInfo.userType != commonType.UserTypeUser {
		ctx.JSON(http.StatusForbidden, responses.ApiResponse{
			Status:  false,
			Message: "Only user can update account!",
		})
		return
	}

	serviceResponse := handler.accountService.Update(uint(id), userInfo.userId, payload)
	if isErrorResponse(ctx, serviceResponse) {
		return
	}

	ctx.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: serviceResponse.Data,
	})
}

func (handler *Account) Delete(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: "Invalid account id!",
		})
		return
	}

	userInfo, ok := getUserInfo(ctx)
	if !ok {
		return
	}

	if userInfo.userType != commonType.UserTypeUser {
		ctx.JSON(http.StatusForbidden, responses.ApiResponse{
			Status:  false,
			Message: "Only user can update account!",
		})
		return
	}

	serviceResponse := handler.accountService.Delete(uint(id), userInfo.userId)
	if isErrorResponse(ctx, serviceResponse) {
		return
	}

	ctx.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: serviceResponse.Data,
	})
}

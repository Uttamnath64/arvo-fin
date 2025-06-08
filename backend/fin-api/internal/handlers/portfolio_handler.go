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

type Portfolio struct {
	container        *storage.Container
	portfolioService *services.Portfolio
}

func NewPortfolio(container *storage.Container) *Portfolio {
	return &Portfolio{
		container:        container,
		portfolioService: services.NewPortfolio(container),
	}
}

func (handler *Portfolio) GetList(ctx *gin.Context) {

	userInfo, ok := getUserInfo(ctx)
	if !ok {
		return
	}

	serviceResponse := handler.portfolioService.GetList(userInfo.userId, userInfo.userType)
	if isErrorResponse(ctx, serviceResponse) {
		return
	}

	portfolioResponse, _ := serviceResponse.Data.(*[]responses.PortfolioResponse)
	ctx.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: portfolioResponse,
	})
}

func (handler *Portfolio) Get(ctx *gin.Context) {

	userInfo, ok := getUserInfo(ctx)
	if !ok {
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: "Invalid portfolio id!",
		})
		return
	}

	serviceResponse := handler.portfolioService.Get(uint(id), userInfo.userId, userInfo.userType)
	if isErrorResponse(ctx, serviceResponse) {
		return
	}

	portfolioResponse, _ := serviceResponse.Data.(*responses.PortfolioResponse)
	ctx.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: portfolioResponse,
	})
}

func (handler *Portfolio) Add(ctx *gin.Context) {

	var payload requests.PortfolioRequest
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
			Message: "Only users can add portfolios!",
		})
	}

	serviceResponse := handler.portfolioService.Add(payload, userInfo.userId)
	if isErrorResponse(ctx, serviceResponse) {
		return
	}

	ctx.JSON(http.StatusOK, responses.ApiResponse{
		Status:  true,
		Message: serviceResponse.Message,
	})
}

func (handler *Portfolio) Update(ctx *gin.Context) {

	var payload requests.PortfolioRequest
	if !bindAndValidateJson(ctx, &payload) {
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: "Invalid portfolio id!",
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
			Message: "You are not allowed to update this portfolio!",
		})
	}

	serviceResponse := handler.portfolioService.Update(uint(id), userInfo.userId, payload)
	if isErrorResponse(ctx, serviceResponse) {
		return
	}

	ctx.JSON(http.StatusOK, responses.ApiResponse{
		Status:  true,
		Message: serviceResponse.Message,
	})
}

func (handler *Portfolio) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: "Invalid portfolio id!",
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
			Message: "You are not authorized to delete a portfolio!",
		})
	}

	serviceResponse := handler.portfolioService.Delete(uint(id), userInfo.userId)
	if isErrorResponse(ctx, serviceResponse) {
		return
	}

	ctx.JSON(http.StatusOK, responses.ApiResponse{
		Status:  true,
		Message: serviceResponse.Message,
	})
}

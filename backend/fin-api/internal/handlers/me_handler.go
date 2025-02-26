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

type Me struct {
	container   *storage.Container
	userService *services.User
}

func NewMe(container *storage.Container) *Me {
	return &Me{
		container:   container,
		userService: services.NewUser(container),
	}
}

func (handler *Me) Get(ctx *gin.Context) {

	userInfo, ok := getUserInfo(ctx)
	if !ok {
		return
	}
	userId := userInfo.userId

	// User is admin then
	if userInfo.userType == commonType.Admin {
		id, err := strconv.Atoi(ctx.Param("user_id"))
		if err != nil || id <= 0 {
			ctx.JSON(http.StatusBadRequest, responses.ApiResponse{
				Status:  false,
				Message: "Invalid user id!",
			})
			return
		}
		userId = uint(id)
	}

	serviceResponse := handler.userService.Get(userId)
	if isErrorResponse(ctx, serviceResponse) {
		return
	}

	portfolioResponse, _ := serviceResponse.Data.(*responses.MeResponse)
	ctx.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: portfolioResponse,
	})
}

func (handler *Me) GetSettings(ctx *gin.Context) {

	userInfo, ok := getUserInfo(ctx)
	if !ok {
		return
	}
	userId := userInfo.userId

	// User is admin then
	if userInfo.userType == commonType.Admin {
		id, err := strconv.Atoi(ctx.Param("user_id"))
		if err != nil || id <= 0 {
			ctx.JSON(http.StatusBadRequest, responses.ApiResponse{
				Status:  false,
				Message: "Invalid user id!",
			})
			return
		}
		userId = uint(id)
	}

	serviceResponse := handler.userService.GetSettings(userId)
	if isErrorResponse(ctx, serviceResponse) {
		return
	}

	portfolioResponse, _ := serviceResponse.Data.(*responses.MeResponse)
	ctx.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: portfolioResponse,
	})
}

func (handler *Me) Update(ctx *gin.Context) {

	var payload requests.MeRequest
	if !bindAndValidateJson(ctx, &payload) {
		return
	}

	userInfo, ok := getUserInfo(ctx)
	if !ok {
		return
	}

	if userInfo.userType != commonType.User {
		ctx.JSON(http.StatusForbidden, responses.ApiResponse{
			Status:  false,
			Message: "Only users can add portfolios!",
		})
	}

	serviceResponse := handler.userService.Update(payload, userInfo.userId)
	if isErrorResponse(ctx, serviceResponse) {
		return
	}

	ctx.JSON(http.StatusOK, responses.ApiResponse{
		Status:  true,
		Message: serviceResponse.Message,
	})
}

func (handler *Me) UpdateSettings(ctx *gin.Context) {

	var payload requests.SettingsRequest
	if !bindAndValidateJson(ctx, &payload) {
		return
	}

	userInfo, ok := getUserInfo(ctx)
	if !ok {
		return
	}

	if userInfo.userType != commonType.User {
		ctx.JSON(http.StatusForbidden, responses.ApiResponse{
			Status:  false,
			Message: "Only users can add portfolios!",
		})
	}

	serviceResponse := handler.userService.UpdateSettings(payload, userInfo.userId)
	if isErrorResponse(ctx, serviceResponse) {
		return
	}

	ctx.JSON(http.StatusOK, responses.ApiResponse{
		Status:  true,
		Message: serviceResponse.Message,
	})
}

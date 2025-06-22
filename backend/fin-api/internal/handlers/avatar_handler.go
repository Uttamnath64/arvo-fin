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

type Avatar struct {
	container     *storage.Container
	avatarService *services.Avatar
}

func NewAvatar(container *storage.Container) *Avatar {
	return &Avatar{
		container:     container,
		avatarService: services.NewAvatar(container),
	}
}

func (handler *Avatar) Get(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: "Invalid avatar id!",
		})
		return
	}

	serviceResponse := handler.avatarService.Get(uint(id))
	if isErrorResponse(ctx, serviceResponse) {
		return
	}

	ctx.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: serviceResponse.Data,
	})
}

func (handler *Avatar) GetAvatarsByType(ctx *gin.Context) {

	typeInt, err := strconv.Atoi(ctx.Param("type"))
	if err != nil || typeInt <= 0 {
		ctx.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: "Invalid type!",
		})
		return
	}

	avatarType := commonType.AvatarType(typeInt)
	if !avatarType.IsValid() {
		ctx.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: "Invalid avatar type!",
		})
		return
	}

	serviceResponse := handler.avatarService.GetAvatarsByType(avatarType)
	if isErrorResponse(ctx, serviceResponse) {
		return
	}

	ctx.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: serviceResponse.Data,
	})
}

func (handler *Avatar) Create(ctx *gin.Context) {

	var payload requests.AvatarRequest
	if !bindAndValidateJson(ctx, &payload) {
		return
	}

	userInfo, ok := getUserInfo(ctx)
	if !ok {
		return
	}

	if userInfo.userType != commonType.UserTypeAdmin {
		ctx.JSON(http.StatusForbidden, responses.ApiResponse{
			Status:  false,
			Message: "Only admin can add avatar!",
		})
		return
	}

	serviceResponse := handler.avatarService.Create(payload)
	if isErrorResponse(ctx, serviceResponse) {
		return
	}

	ctx.JSON(http.StatusOK, responses.ApiResponse{
		Status:  true,
		Message: serviceResponse.Message,
	})
}

func (handler *Avatar) Update(ctx *gin.Context) {

	var payload requests.AvatarRequest
	if !bindAndValidateJson(ctx, &payload) {
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: "Invalid avatar id!",
		})
		return
	}

	userInfo, ok := getUserInfo(ctx)
	if !ok {
		return
	}

	if userInfo.userType != commonType.UserTypeAdmin {
		ctx.JSON(http.StatusForbidden, responses.ApiResponse{
			Status:  false,
			Message: "Only admin can update avatar!",
		})
		return
	}

	serviceResponse := handler.avatarService.Update(uint(id), payload)
	if isErrorResponse(ctx, serviceResponse) {
		return
	}

	ctx.JSON(http.StatusOK, responses.ApiResponse{
		Status:  true,
		Message: serviceResponse.Message,
	})
}

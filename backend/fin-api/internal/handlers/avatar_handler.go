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

func (handler *Avatar) Get(c *gin.Context) {

	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: "Invalid avatar id!",
		})
		return
	}

	serviceResponse := handler.avatarService.Get(rctx, uint(id))
	if isErrorResponse(c, serviceResponse) {
		return
	}

	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: serviceResponse.Data,
	})
}

func (handler *Avatar) GetAvatarsByType(c *gin.Context) {

	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}

	typeInt, err := strconv.Atoi(c.Param("type"))
	if err != nil || typeInt <= 0 {
		c.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: "Invalid type!",
		})
		return
	}

	avatarType := commonType.AvatarType(typeInt)
	if !avatarType.IsValid() {
		c.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: "Invalid avatar type!",
		})
		return
	}

	serviceResponse := handler.avatarService.GetAvatarsByType(rctx, avatarType)
	if isErrorResponse(c, serviceResponse) {
		return
	}

	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: serviceResponse.Data,
	})
}

func (handler *Avatar) Create(c *gin.Context) {

	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}

	var payload requests.AvatarRequest
	if !bindAndValidateJson(c, &payload) {
		return
	}

	if rctx.UserType != commonType.UserTypeAdmin {
		c.JSON(http.StatusForbidden, responses.ApiResponse{
			Status:  false,
			Message: "Only admin can add avatar!",
		})
		return
	}

	serviceResponse := handler.avatarService.Create(rctx, payload)
	if isErrorResponse(c, serviceResponse) {
		return
	}

	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:  true,
		Message: serviceResponse.Message,
	})
}

func (handler *Avatar) Update(c *gin.Context) {

	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}

	var payload requests.AvatarRequest
	if !bindAndValidateJson(c, &payload) {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: "Invalid avatar id!",
		})
		return
	}

	if rctx.UserType != commonType.UserTypeAdmin {
		c.JSON(http.StatusForbidden, responses.ApiResponse{
			Status:  false,
			Message: "Only admin can update avatar!",
		})
		return
	}

	serviceResponse := handler.avatarService.Update(rctx, uint(id), payload)
	if isErrorResponse(c, serviceResponse) {
		return
	}

	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:  true,
		Message: serviceResponse.Message,
	})
}

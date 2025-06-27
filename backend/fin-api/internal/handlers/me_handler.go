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

func (handler *Me) Get(c *gin.Context) {

	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}
	userId := rctx.UserID

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

	serviceResponse := handler.userService.Get(rctx, userId)
	if isErrorResponse(c, serviceResponse) {
		return
	}

	response, _ := serviceResponse.Data.(*responses.MeResponse)
	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: response,
	})
}

func (handler *Me) GetSettings(c *gin.Context) {

	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}
	userId := rctx.UserID

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

	serviceResponse := handler.userService.GetSettings(rctx, userId)
	if isErrorResponse(c, serviceResponse) {
		return
	}

	response, _ := serviceResponse.Data.(*responses.SettingsResponse)
	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: response,
	})
}

func (handler *Me) Update(c *gin.Context) {

	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}

	var payload requests.MeRequest
	if !bindAndValidateJson(c, &payload) {
		return
	}

	if rctx.UserType != commonType.UserTypeUser {
		c.JSON(http.StatusForbidden, responses.ApiResponse{
			Status:  false,
			Message: "Only users can update profile!",
		})
	}

	serviceResponse := handler.userService.Update(rctx, payload, rctx.UserID)
	if isErrorResponse(c, serviceResponse) {
		return
	}

	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:  true,
		Message: serviceResponse.Message,
	})
}

func (handler *Me) UpdateSettings(c *gin.Context) {

	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}

	var payload requests.SettingsRequest
	if !bindAndValidateJson(c, &payload) {
		return
	}

	if rctx.UserType != commonType.UserTypeUser {
		c.JSON(http.StatusForbidden, responses.ApiResponse{
			Status:  false,
			Message: "Only users can update setting!",
		})
	}

	serviceResponse := handler.userService.UpdateSettings(rctx, payload, rctx.UserID)
	if isErrorResponse(c, serviceResponse) {
		return
	}

	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:  true,
		Message: serviceResponse.Message,
	})
}

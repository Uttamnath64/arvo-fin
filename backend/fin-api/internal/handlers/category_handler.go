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

type Category struct {
	container *storage.Container
	service   *services.Category
}

func NewCategory(container *storage.Container) *Category {
	return &Category{
		container: container,
		service:   services.NewCategory(container),
	}
}

func (handler *Category) Get(c *gin.Context) {

	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: "Invalid category id!",
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

func (handler *Category) GetList(c *gin.Context) {

	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}
	userId := rctx.UserID

	portfolioId, err := strconv.Atoi(c.Param("portfolioId"))
	if err != nil || portfolioId <= 0 {
		c.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: "Invalid category id!",
		})
		return
	}

	// User is admin then
	if rctx.UserType == commonType.UserTypeAdmin {
		userIdStr := c.Query("userId")
		if userIdStr != "" {
			userIdInt, err := strconv.Atoi(userIdStr)
			if err != nil || userIdInt <= 0 {
				c.JSON(http.StatusBadRequest, responses.ApiResponse{
					Status:  false,
					Message: "Invalid userId!",
				})
				return
			}
			userId = uint(userIdInt)
		} else {
			c.JSON(http.StatusBadRequest, responses.ApiResponse{
				Status:  false,
				Message: "userId query param is required for admin users!",
			})
			return
		}
	}

	serviceResponse := handler.service.GetList(rctx, uint(portfolioId), userId)
	if isErrorResponse(c, serviceResponse) {
		return
	}

	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: serviceResponse.Data,
	})
}

func (handler *Category) Create(c *gin.Context) {

	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}

	var payload requests.CategoryRequest
	if !bindAndValidateJson(c, &payload) {
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

func (handler *Category) Update(c *gin.Context) {

	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}

	var payload requests.CategoryRequest
	if !bindAndValidateJson(c, &payload) {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: "Invalid category id!",
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

func (handler *Category) Delete(c *gin.Context) {

	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}

	portfolioId, err := strconv.Atoi(c.Param("portfolioId"))
	if err != nil || portfolioId <= 0 {
		c.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: "Invalid category id!",
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: "Invalid category id!",
		})
		return
	}

	serviceResponse := handler.service.Delete(rctx, uint(portfolioId), uint(id))
	if isErrorResponse(c, serviceResponse) {
		return
	}

	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: serviceResponse.Data,
	})
}

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

func (handler *Portfolio) GetList(c *gin.Context) {

	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}
	userId := rctx.UserID

	if rctx.UserType == commonType.UserTypeAdmin {
		userIdInt, err := strconv.Atoi(c.Param("userId"))
		if err != nil || userIdInt <= 0 {
			c.JSON(http.StatusBadRequest, responses.ApiResponse{
				Status:  false,
				Message: "Invalid user id!",
			})
			return
		}
		userId = uint(userIdInt)
	}

	serviceResponse := handler.portfolioService.GetList(rctx, userId)
	if isErrorResponse(c, serviceResponse) {
		return
	}

	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: serviceResponse.Data,
	})
}

func (handler *Portfolio) Get(c *gin.Context) {

	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: "Invalid portfolio id!",
		})
		return
	}

	serviceResponse := handler.portfolioService.Get(rctx, uint(id), rctx.UserID, rctx.UserType)
	if isErrorResponse(c, serviceResponse) {
		return
	}

	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: serviceResponse.Data,
	})
}

func (handler *Portfolio) Create(c *gin.Context) {

	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}

	var payload requests.PortfolioRequest
	if !bindAndValidateJson(c, &payload) {
		return
	}

	if rctx.UserType != commonType.UserTypeUser {
		c.JSON(http.StatusForbidden, responses.ApiResponse{
			Status:  false,
			Message: "Only users can add portfolios!",
		})
	}

	serviceResponse := handler.portfolioService.Create(rctx, payload, rctx.UserID)
	if isErrorResponse(c, serviceResponse) {
		return
	}

	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:  true,
		Message: serviceResponse.Message,
	})
}

func (handler *Portfolio) Update(c *gin.Context) {

	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}

	var payload requests.PortfolioRequest
	if !bindAndValidateJson(c, &payload) {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: "Invalid portfolio id!",
		})
		return
	}

	if rctx.UserType != commonType.UserTypeUser {
		c.JSON(http.StatusForbidden, responses.ApiResponse{
			Status:  false,
			Message: "You are not allowed to update this portfolio!",
		})
	}

	serviceResponse := handler.portfolioService.Update(rctx, uint(id), rctx.UserID, payload)
	if isErrorResponse(c, serviceResponse) {
		return
	}

	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:  true,
		Message: serviceResponse.Message,
	})
}

func (handler *Portfolio) Delete(c *gin.Context) {

	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: "Invalid portfolio id!",
		})
		return
	}

	if rctx.UserType != commonType.UserTypeUser {
		c.JSON(http.StatusForbidden, responses.ApiResponse{
			Status:  false,
			Message: "You are not authorized to delete a portfolio!",
		})
	}

	serviceResponse := handler.portfolioService.Delete(rctx, uint(id), rctx.UserID)
	if isErrorResponse(c, serviceResponse) {
		return
	}

	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:  true,
		Message: serviceResponse.Message,
	})
}

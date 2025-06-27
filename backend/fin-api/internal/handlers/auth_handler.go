package handlers

import (
	"net/http"

	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/responses"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"github.com/Uttamnath64/arvo-fin/fin-api/internal/services"
	"github.com/gin-gonic/gin"
)

type Auth struct {
	container   *storage.Container
	authService *services.Auth
}

func NewAuth(container *storage.Container) *Auth {
	return &Auth{
		container:   container,
		authService: services.NewAuth(container),
	}
}

func (handler *Auth) Login(c *gin.Context) {

	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}

	var payload requests.LoginRequest
	if !bindAndValidateJson(c, &payload) {
		return
	}

	serviceResponse := handler.authService.Login(rctx, payload, c.Request.UserAgent(), c.ClientIP())
	if isErrorResponse(c, serviceResponse) {
		return
	}

	authR, _ := serviceResponse.Data.(responses.AuthResponse)
	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: authR,
	})
}

func (handler *Auth) Register(c *gin.Context) {

	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}

	var payload requests.RegisterRequest
	if !bindAndValidateJson(c, &payload) {
		return
	}

	serviceResponse := handler.authService.Register(rctx, payload, c.Request.UserAgent(), c.ClientIP())
	if isErrorResponse(c, serviceResponse) {
		return
	}

	authR, _ := serviceResponse.Data.(responses.AuthResponse)
	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: authR,
	})
}

func (handler *Auth) SendOTP(c *gin.Context) {

	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}

	var payload requests.SentOTPRequest
	if !bindAndValidateJson(c, &payload) {
		return
	}

	serviceResponse := handler.authService.SendOTP(rctx, payload)
	if isErrorResponse(c, serviceResponse) {
		return
	}

	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:  true,
		Message: "OTP sent successfully to the email address!",
	})
}

func (handler *Auth) ResetPassword(c *gin.Context) {

	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}

	var payload requests.ResetPasswordRequest
	if !bindAndValidateJson(c, &payload) {
		return
	}

	// Reset password
	serviceResponse := handler.authService.ResetPassword(rctx, payload, c.Request.UserAgent(), c.ClientIP())
	if isErrorResponse(c, serviceResponse) {
		return
	}

	authR, _ := serviceResponse.Data.(responses.AuthResponse)
	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: authR,
	})
}

func (handler *Auth) Token(c *gin.Context) {

	rctx, ok := getRequestContext(c)
	if !ok {
		return
	}

	var payload requests.TokenRequest
	if !bindAndValidateJson(c, &payload) {
		return
	}

	// Get token
	serviceResponse := handler.authService.GetToken(rctx, payload, c.Request.UserAgent(), c.ClientIP())
	if isErrorResponse(c, serviceResponse) {
		return
	}

	authR, _ := serviceResponse.Data.(responses.AuthResponse)
	c.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: authR,
	})

}

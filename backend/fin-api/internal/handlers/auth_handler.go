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

func (handler *Auth) Login(ctx *gin.Context) {

	var payload requests.LoginRequest
	if !bindAndValidateJson(ctx, &payload) {
		return
	}

	serviceResponse := handler.authService.Login(payload, ctx.Request.UserAgent(), ctx.ClientIP())
	if isErrorResponse(ctx, serviceResponse) {
		return
	}

	authR, _ := serviceResponse.Data.(responses.AuthResponse)
	ctx.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: authR,
	})
}

func (handler *Auth) Register(ctx *gin.Context) {

	var payload requests.RegisterRequest
	if !bindAndValidateJson(ctx, &payload) {
		return
	}

	serviceResponse := handler.authService.Register(payload, ctx.Request.UserAgent(), ctx.ClientIP())
	if isErrorResponse(ctx, serviceResponse) {
		return
	}

	authR, _ := serviceResponse.Data.(responses.AuthResponse)
	ctx.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: authR,
	})
}

func (handler *Auth) SendOTP(ctx *gin.Context) {

	var payload requests.SentOTPRequest
	if !bindAndValidateJson(ctx, &payload) {
		return
	}

	serviceResponse := handler.authService.SendOTP(payload)
	if isErrorResponse(ctx, serviceResponse) {
		return
	}

	ctx.JSON(http.StatusOK, responses.ApiResponse{
		Status:  true,
		Message: "OTP sent successfully to the email address!",
	})
}

func (handler *Auth) ResetPassword(ctx *gin.Context) {

	var payload requests.ResetPasswordRequest
	if !bindAndValidateJson(ctx, &payload) {
		return
	}

	// Reset password
	serviceResponse := handler.authService.ResetPassword(payload, ctx.Request.UserAgent(), ctx.ClientIP())
	if isErrorResponse(ctx, serviceResponse) {
		return
	}

	authR, _ := serviceResponse.Data.(responses.AuthResponse)
	ctx.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: authR,
	})
}

func (handler *Auth) Token(ctx *gin.Context) {

	var payload requests.TokenRequest
	if !bindAndValidateJson(ctx, &payload) {
		return
	}

	// Get token
	serviceResponse := handler.authService.GetToken(payload, ctx.Request.UserAgent(), ctx.ClientIP())
	if isErrorResponse(ctx, serviceResponse) {
		return
	}

	authR, _ := serviceResponse.Data.(responses.AuthResponse)
	ctx.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: authR,
	})

}

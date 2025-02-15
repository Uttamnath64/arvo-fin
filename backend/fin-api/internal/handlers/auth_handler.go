package handlers

import (
	"net/http"

	"github.com/Uttamnath64/arvo-fin/app/common"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/responses"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"github.com/Uttamnath64/arvo-fin/fin-api/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	container   *storage.Container
	authService *services.AuthService
}

func NewAuthHandler(container *storage.Container) *AuthHandler {
	return &AuthHandler{
		container:   container,
		authService: services.NewAuthService(container),
	}
}

func (handler *AuthHandler) LoginHandler(ctx *gin.Context) {
	var payload requests.LoginRequest

	// Bind and validate input
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: "Invalid request payload. Please check the input data format!",
			Details: err.Error(),
		})
		return
	}

	if err := payload.IsValid(); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	// Login
	serviceResponse := handler.authService.Login(payload, ctx.ClientIP())

	if serviceResponse.HasError() {
		switch serviceResponse.StatusCode {
		case common.StatusNotFound:
			ctx.JSON(http.StatusBadRequest, responses.ApiResponse{
				Status:  false,
				Message: serviceResponse.Message,
				Details: serviceResponse.Error.Error(),
			})
		case common.StatusValidationError:
			ctx.JSON(http.StatusUnauthorized, responses.ApiResponse{
				Status:  false,
				Message: serviceResponse.Message,
				Details: serviceResponse.Error.Error(),
			})
		case common.StatusServerError:
			ctx.JSON(http.StatusInternalServerError, responses.ApiResponse{
				Status:  false,
				Message: serviceResponse.Message,
				Details: serviceResponse.Error.Error(),
			})
		}
		return
	}

	authR, _ := serviceResponse.Data.(responses.AuthResponse)

	ctx.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: authR,
	})
}

func (handler *AuthHandler) RegisterHandler(ctx *gin.Context) {
	var payload requests.RegisterRequest

	// Bind and validate input
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: "Invalid request payload. Please check the input data format!",
			Details: err.Error(),
		})
		return
	}

	if err := payload.IsValid(); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	// Register
	serviceResponse := handler.authService.Register(payload, ctx.ClientIP())

	if serviceResponse.HasError() {
		switch serviceResponse.StatusCode {
		case common.StatusNotFound:
			ctx.JSON(http.StatusBadRequest, responses.ApiResponse{
				Status:  false,
				Message: serviceResponse.Message,
				Details: serviceResponse.Error.Error(),
			})
		case common.StatusValidationError:
			ctx.JSON(http.StatusUnauthorized, responses.ApiResponse{
				Status:  false,
				Message: serviceResponse.Message,
				Details: serviceResponse.Error.Error(),
			})
		case common.StatusServerError:
			ctx.JSON(http.StatusInternalServerError, responses.ApiResponse{
				Status:  false,
				Message: serviceResponse.Message,
				Details: serviceResponse.Error.Error(),
			})
		}
		return
	}

	authR, _ := serviceResponse.Data.(responses.AuthResponse)

	ctx.JSON(http.StatusOK, responses.ApiResponse{
		Status:   true,
		Message:  serviceResponse.Message,
		Metadata: authR,
	})
}

func (handler *AuthHandler) SentOTPHandler(ctx *gin.Context) {
	var payload requests.SentOTPRequest

	// Bind and validate input
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: "Invalid request payload. Please check the input data format!",
			Details: err.Error(),
		})
		return
	}

	if err := payload.IsValid(); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	// Send OTP to email
	serviceResponse := handler.authService.SentOTP(payload)

	if serviceResponse.HasError() {
		switch serviceResponse.StatusCode {
		case common.StatusNotFound:
			ctx.JSON(http.StatusBadRequest, responses.ApiResponse{
				Status:  false,
				Message: serviceResponse.Message,
				Details: serviceResponse.Error.Error(),
			})
		case common.StatusValidationError:
			ctx.JSON(http.StatusUnauthorized, responses.ApiResponse{
				Status:  false,
				Message: serviceResponse.Message,
				Details: serviceResponse.Error.Error(),
			})
		case common.StatusServerError:
			ctx.JSON(http.StatusInternalServerError, responses.ApiResponse{
				Status:  false,
				Message: serviceResponse.Message,
				Details: serviceResponse.Error.Error(),
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, responses.ApiResponse{
		Status:  true,
		Message: "OTP sent successfully to the email address!",
	})
}

func (handler *AuthHandler) ResetPasswordHandler(ctx *gin.Context) {
	var payload requests.ResetPasswordRequest

	// Bind and validate input
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: "Invalid request payload. Please check the input data format!",
			Details: err.Error(),
		})
		return
	}

	if err := payload.IsValid(); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ApiResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	// Reset password
	serviceResponse := handler.authService.ResetPassword(payload)

	if serviceResponse.HasError() {
		switch serviceResponse.StatusCode {
		case common.StatusNotFound:
			ctx.JSON(http.StatusBadRequest, responses.ApiResponse{
				Status:  false,
				Message: serviceResponse.Message,
				Details: serviceResponse.Error.Error(),
			})
		case common.StatusValidationError:
			ctx.JSON(http.StatusUnauthorized, responses.ApiResponse{
				Status:  false,
				Message: serviceResponse.Message,
				Details: serviceResponse.Error.Error(),
			})
		case common.StatusServerError:
			ctx.JSON(http.StatusInternalServerError, responses.ApiResponse{
				Status:  false,
				Message: serviceResponse.Message,
				Details: serviceResponse.Error.Error(),
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, responses.ApiResponse{
		Status:  true,
		Message: "Password reset successfully!",
	})
}

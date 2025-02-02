package handlers

import (
	"net/http"

	"github.com/Uttamnath64/arvo-fin/app/common"
	"github.com/Uttamnath64/arvo-fin/app/config"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/responses"
	"github.com/Uttamnath64/arvo-fin/fin-api/internal/services"
	"github.com/Uttamnath64/arvo-fin/pkg/logger"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	config      *config.Config
	logger      *logger.Logger
	authService *services.AuthService
	env         *config.AppEnv
}

func NewAuthHandler(config *config.Config, logger *logger.Logger, env *config.AppEnv) *AuthHandler {
	return &AuthHandler{
		config:      config,
		logger:      logger,
		authService: services.NewAuthService(config, logger, env),
		env:         env,
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

	// Find user by email
	serviceResponse := handler.authService.Login(ctx.ClientIP(), payload)

	if serviceResponse.HasError() {
		switch serviceResponse.StatusCode {
		case common.StatusNotFound:
			ctx.JSON(http.StatusBadRequest, responses.ApiResponse{
				Status:  false,
				Message: serviceResponse.Error.Error(),
			})
		case common.StatusValidationError:
			ctx.JSON(http.StatusUnauthorized, responses.ApiResponse{
				Status:  false,
				Message: serviceResponse.Error.Error(),
			})
		case common.StatusServerError:
			ctx.JSON(http.StatusInternalServerError, responses.ApiResponse{
				Status:  false,
				Message: serviceResponse.Error.Error(),
			})
		}
		return
	}

	authR, _ := serviceResponse.Data.(responses.AuthResponse)

	ctx.JSON(http.StatusOK, responses.ApiResponse{
		Status:   false,
		Message:  serviceResponse.Message,
		Metadata: authR,
	})
}

func (handler *AuthHandler) RegisterHandler(c *gin.Context) {

}

func (handler *AuthHandler) GetOTPHandler(c *gin.Context) {

}

func (handler *AuthHandler) ForgotPasswordHandler(c *gin.Context) {

}

func (handler *AuthHandler) ResetPasswordHandler(c *gin.Context) {

}

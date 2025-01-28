package handlers

import (
	"net/http"

	"github.com/Uttamnath64/arvo-fin/app/config"
	"github.com/Uttamnath64/arvo-fin/fin-api/internal/servers/requests"
	response "github.com/Uttamnath64/arvo-fin/fin-api/internal/servers/responses"
	"github.com/Uttamnath64/arvo-fin/fin-api/internal/services"
	"github.com/Uttamnath64/arvo-fin/pkg/logger"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	config      *config.Config
	logger      *logger.Logger
	authService *services.AuthService
}

func NewAuthHandler(config *config.Config, logger *logger.Logger) *AuthHandler {
	return &AuthHandler{
		config:      config,
		logger:      logger,
		authService: services.NewAuthService(config, logger),
	}
}

func (handler *AuthHandler) LoginHandler(c *gin.Context) {
	var payload requests.LoginRequest

	// Bind and validate input
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  false,
			Message: "Invalid input!",
			Details: err.Error(),
		})
		return
	}

	if err := payload.IsValid(); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	if err := payload.IsValid(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	// Find user by email
	user, err := handler.authService.FindUserByUsernameEmail(req.UsernameEmail)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse{
			Status:  false,
			Message: "Invalid email or password!",
			Details: err.Error(),
		})
		return
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Invalid email or password",
		})
		return
	}

}

func (handler *AuthHandler) RegisterHandler(c *gin.Context) {

}

func (handler *AuthHandler) GetOTPHandler(c *gin.Context) {

}

func (handler *AuthHandler) ForgotPasswordHandler(c *gin.Context) {

}

func (handler *AuthHandler) ResetPasswordHandler(c *gin.Context) {

}

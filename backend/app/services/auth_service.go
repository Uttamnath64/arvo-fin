package services

import (
	"errors"

	"github.com/Uttamnath64/arvo-fin/app/auth"
	"github.com/Uttamnath64/arvo-fin/app/common"
	"github.com/Uttamnath64/arvo-fin/app/config"
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/repository"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/responses"
	"github.com/Uttamnath64/arvo-fin/pkg/logger"
	"gorm.io/gorm"
)

type AuthService struct {
	config   *config.Config
	logger   *logger.Logger
	env      *config.AppEnv
	userRepo *repository.UserRepository
}

func NewAuthService(config *config.Config, logger *logger.Logger, env *config.AppEnv) *AuthService {
	return &AuthService{
		config:   config,
		logger:   logger,
		userRepo: repository.NewUserRepository(config, logger),
		env:      env,
	}
}

func (service *AuthService) Login(ip string, payload requests.LoginRequest) (res *responses.ServiceResponse) {
	var user models.User

	// Check user
	err := service.userRepo.GetUser(payload.UsernameEmail, &user)
	if err != gorm.ErrRecordNotFound {
		res = &responses.ServiceResponse{
			StatusCode: common.StatusNotFound,
			Message:    "Username/Email not found!",
			Error:      errors.New("Authentication failed!"),
		}
		return
	}

	// Validate password
	if err := Validate.VerifyPassword(user.Password, payload.Password); err != nil {
		res = &responses.ServiceResponse{
			StatusCode: common.StatusValidationError,
			Message:    "Invalid password!",
			Error:      errors.New("Authentication failed!"),
		}
		return
	}

	// Create Token
	authRepo := repository.NewAuthRepository(service.config, service.logger)
	authHeler := auth.New(service.config, service.env, service.logger, authRepo)
	accessToken, refreshToken, err := authHeler.GenerateToken(user.ID, common.USER_TYPE_USER, ip)
	if err != nil {
		res = &responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      errors.New("Something went wrong!"),
		}
		return
	}

	// Response
	res = &responses.ServiceResponse{
		StatusCode: common.StatusSuccess,
		Message:    "User login successfully!",
		Data: responses.AuthResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}
	return
}

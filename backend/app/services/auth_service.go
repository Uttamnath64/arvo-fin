package services

import (
	"errors"

	"github.com/Uttamnath64/arvo-fin/app/auth"
	"github.com/Uttamnath64/arvo-fin/app/common"
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/repository"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/responses"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"gorm.io/gorm"
)

type AuthService struct {
	container *storage.Container
	userRepo  *repository.UserRepository
}

func NewAuthService(container *storage.Container) *AuthService {
	return &AuthService{
		container: container,
		userRepo:  repository.NewUserRepository(container),
	}
}

func (service *AuthService) Login(payload requests.LoginRequest, ip string) responses.ServiceResponse {
	var user models.User

	// Check user
	err := service.userRepo.GetUser(payload.UsernameEmail, &user)
	if err == gorm.ErrRecordNotFound {
		return responses.ServiceResponse{
			StatusCode: common.StatusNotFound,
			Message:    "Username/Email not found!",
			Error:      errors.New("Authentication failed!"),
		}
	}

	// Database issue
	if err != nil {
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      errors.New("Something went wrong!"),
		}
	}

	// Validate password
	if err := Validate.VerifyPassword(user.Password, payload.Password); err != nil {
		return responses.ServiceResponse{
			StatusCode: common.StatusValidationError,
			Message:    "Invalid password!",
			Error:      errors.New("Authentication failed!"),
		}
	}

	// Create Token
	authRepo := repository.NewAuthRepository(service.container)
	authHeler := auth.New(service.container, authRepo)
	accessToken, refreshToken, err := authHeler.GenerateToken(user.ID, common.USER_TYPE_USER, ip)
	if err != nil {
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      errors.New("Something went wrong!"),
		}
	}

	// Response
	return responses.ServiceResponse{
		StatusCode: common.StatusSuccess,
		Message:    "User login successfully!",
		Data: responses.AuthResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}
}

func (service *AuthService) Register(payload requests.RegisterRequest, ip string) responses.ServiceResponse {
	var (
		user     models.User
		err      error
		isExists bool
	)

	// Check username
	isExists, err = service.userRepo.UsernameExists(payload.Username)
	if err != nil {
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      errors.New("Something went wrong!"),
		}
	}
	if isExists {
		return responses.ServiceResponse{
			StatusCode: common.StatusValidationError,
			Message:    "Username already exists!",
			Error:      errors.New("Username already exists!"),
		}
	}

	// Check email
	isExists, err = service.userRepo.EmailExists(payload.Email)
	if err != nil {
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      errors.New("Something went wrong!"),
		}
	}
	if isExists {
		return responses.ServiceResponse{
			StatusCode: common.StatusValidationError,
			Message:    "Email already exists!",
			Error:      errors.New("Email already exists!"),
		}
	}

	// Verify OTP
	otpService := NewOTPService(service.container.Redis, 250)
	err = otpService.VerifyOTP(payload.Email, payload.OTP)
	if err != nil {
		return responses.ServiceResponse{
			StatusCode: common.StatusValidationError,
			Message:    "Invalid OTP!",
			Error:      errors.New("Invalid OTP!"),
		}
	}

	// Hash password
	user.Password, err = Validate.HashPassword(user.Password)
	if err != nil {
		return responses.ServiceResponse{
			StatusCode: common.StatusValidationError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      errors.New("Something went wrong!"),
		}
	}

	// Create Token
	authRepo := repository.NewAuthRepository(service.container)
	authHeler := auth.New(service.container, authRepo)
	accessToken, refreshToken, err := authHeler.GenerateToken(user.ID, common.USER_TYPE_USER, ip)
	if err != nil {
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      errors.New("Something went wrong!"),
		}
	}

	// Response
	return responses.ServiceResponse{
		StatusCode: common.StatusSuccess,
		Message:    "User login successfully!",
		Data: responses.AuthResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}
}

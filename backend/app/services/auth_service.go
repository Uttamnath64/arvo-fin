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
	"github.com/Uttamnath64/arvo-fin/app/templates"
	"gorm.io/gorm"
)

type AuthService struct {
	container    *storage.Container
	userRepo     *repository.UserRepository
	otpService   *OTPService
	emailService *EmailService
}

func NewAuthService(container *storage.Container) *AuthService {
	return &AuthService{
		container:    container,
		userRepo:     repository.NewUserRepository(container),
		otpService:   NewOTPService(container.Redis, 300),
		emailService: NewEmailService(container),
	}
}

func (service *AuthService) Login(payload requests.LoginRequest, ip string) responses.ServiceResponse {
	var user models.User

	// Check user
	err := service.userRepo.GetUser(payload.UsernameEmail, payload.UsernameEmail, &user)
	if err == gorm.ErrRecordNotFound {
		return responses.ServiceResponse{
			StatusCode: common.StatusNotFound,
			Message:    "Username/Email not found!",
			Error:      err,
		}
	}
	if err != nil {
		service.container.Logger.Error("auth.service.Login-GetUser", err.Error(), payload.UsernameEmail)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      err,
		}
	}

	// Validate password
	if err := Validate.VerifyPassword(user.Password, payload.Password); err != nil {
		return responses.ServiceResponse{
			StatusCode: common.StatusValidationError,
			Message:    "Invalid password!",
			Error:      err,
		}
	}

	// Create Token
	authRepo := repository.NewAuthRepository(service.container)
	authHeler := auth.New(service.container, authRepo)
	accessToken, refreshToken, err := authHeler.GenerateToken(user.ID, common.USER_TYPE_USER, ip)
	if err != nil {
		service.container.Logger.Error("auth.service.Login-GenerateToken", err.Error(), user.ID, common.USER_TYPE_USER, ip)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      err,
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
		err      error
		isExists bool
		password string
	)

	// Check username
	isExists, err = service.userRepo.UsernameExists(payload.Username)
	if err != nil {
		service.container.Logger.Error("auth.service.Register-UsernameExists", err.Error(), payload.Email)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      err,
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
		service.container.Logger.Error("auth.service.Register-EmailExists", err.Error(), payload.Email)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      err,
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
	otpService := NewOTPService(service.container.Redis, 300)
	err = otpService.VerifyOTP(payload.Email, common.Register, payload.OTP)
	if err != nil {
		return responses.ServiceResponse{
			StatusCode: common.StatusValidationError,
			Message:    "Invalid OTP!",
			Error:      errors.New("Invalid OTP!"),
		}
	}

	// Hash password
	password, err = Validate.HashPassword(payload.Password)
	if err != nil {
		service.container.Logger.Error("auth.service.Register-HashPassword", err.Error(), payload.Email, payload.Password)
		return responses.ServiceResponse{
			StatusCode: common.StatusValidationError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      err,
		}
	}

	// Create user
	userId, err := service.userRepo.CreateUser(&models.User{
		Name:         payload.Name,
		Username:     payload.Username,
		Email:        payload.Email,
		MobileNumber: payload.MobileNumber,
		Password:     password,
	})
	if err != nil {
		service.container.Logger.Error("auth.service.Register-CreateUser", err.Error(), payload.Username, payload.Email, password)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      err,
		}
	}

	// Create Token
	authRepo := repository.NewAuthRepository(service.container)
	authHeler := auth.New(service.container, authRepo)
	accessToken, refreshToken, err := authHeler.GenerateToken(userId, common.USER_TYPE_USER, ip)
	if err != nil {
		service.container.Logger.Error("auth.service.Register-GenerateToken", err.Error(), userId, common.USER_TYPE_USER, ip)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      err,
		}
	}

	// Response
	service.container.Logger.Info("auth.service.Register", "User login successfully!", userId, common.USER_TYPE_USER, ip)
	return responses.ServiceResponse{
		StatusCode: common.StatusSuccess,
		Message:    "User login successfully!",
		Data: responses.AuthResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}
}

func (service *AuthService) SentOTP(payload requests.SentOTPRequest) responses.ServiceResponse {

	// Check email
	if payload.Type != common.Register {
		isExists, err := service.userRepo.EmailExists(payload.Email)
		if err != nil {
			service.container.Logger.Error("auth.service.SentOTP-EmailExists", err.Error(), payload.Email)
			return responses.ServiceResponse{
				StatusCode: common.StatusServerError,
				Message:    "Oops! Something went wrong. Please try again later.",
				Error:      err,
			}
		}
		if !isExists {
			return responses.ServiceResponse{
				StatusCode: common.StatusValidationError,
				Message:    "User not found!",
				Error:      errors.New("User not found!"),
			}
		}
	}

	otp := service.otpService.GenerateOTP()
	if err := service.otpService.SaveOTP(payload.Email, payload.Type, otp); err != nil {
		service.container.Logger.Error("auth-service-SentOTP-RedisSaveOTP", err.Error(), payload.Email, otp)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "OTP generation failed!",
			Error:      err,
		}
	}

	// Send OTP to email
	data := map[string]string{
		"OTP":   otp,
		"Email": payload.Email,
	}
	err := service.emailService.SendEmail(payload.Email, "OTP Verification", templates.OTPVerificationEmailTemplate, data, []string{})
	if err != nil {
		service.container.Logger.Error("auth.service.SentOTP-EmailSend", err.Error(), "OTP Verification", templates.OTPVerificationEmailTemplate, data)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Failed to send OTP to email!",
			Error:      err,
		}
	}

	service.container.Logger.Info("auth.service.SentOTP", "OTP sent successfully to the email address!", payload.Email, "OTP Verification", templates.OTPVerificationEmailTemplate, data)
	return responses.ServiceResponse{
		StatusCode: common.StatusSuccess,
		Message:    "OTP sent successfully to the email address!",
	}
}

func (service *AuthService) ResetPassword(payload requests.ResetPasswordRequest) responses.ServiceResponse {
	var user models.User

	// Check user
	err := service.userRepo.GetUser("", payload.Email, &user)
	if err == gorm.ErrRecordNotFound {
		return responses.ServiceResponse{
			StatusCode: common.StatusNotFound,
			Message:    "User not found!",
			Error:      err,
		}
	}
	if err != nil {
		service.container.Logger.Error("auth.service.ResetPassword-GetUser", err.Error(), payload)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      err,
		}
	}

	// Validate password new password
	if err := Validate.VerifyPassword(user.Password, payload.Password); err == nil {
		return responses.ServiceResponse{
			StatusCode: common.StatusValidationError,
			Message:    "Password is the same as the previous one!",
			Error:      err,
		}
	}

	// Verify OTP
	err = service.otpService.VerifyOTP(payload.Email, common.ResetPassword, payload.OTP)
	if err != nil {
		return responses.ServiceResponse{
			StatusCode: common.StatusValidationError,
			Message:    "Invalid OTP!",
			Error:      err,
		}
	}

	// Hash password
	password, err := Validate.HashPassword(payload.Password)
	if err != nil {
		service.container.Logger.Error("auth.service.ResetPassword-HashPassword", err.Error(), payload)
		return responses.ServiceResponse{
			StatusCode: common.StatusValidationError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      err,
		}
	}

	// Update password in user table
	err = service.userRepo.UpdatePasswordByEmail(payload.Email, password)
	if err != nil {
		service.container.Logger.Error("auth.service.ResetPassword-UpdatePasswordByEmail", err.Error(), payload, password)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      err,
		}
	}

	return responses.ServiceResponse{
		StatusCode: common.StatusSuccess,
		Message:    "Password updated successfully!",
	}
}

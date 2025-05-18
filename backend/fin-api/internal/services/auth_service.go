package services

import (
	"errors"

	"github.com/Uttamnath64/arvo-fin/app/auth"
	"github.com/Uttamnath64/arvo-fin/app/common"
	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/repository"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/responses"
	appService "github.com/Uttamnath64/arvo-fin/app/services"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"github.com/Uttamnath64/arvo-fin/app/templates"
	"gorm.io/gorm"
)

type Auth struct {
	container    *storage.Container
	userRepo     *repository.User
	otpService   *appService.OTPService
	emailService *appService.EmailService
}

func NewAuth(container *storage.Container) *Auth {
	return &Auth{
		container:    container,
		userRepo:     repository.NewUser(container),
		otpService:   appService.NewOTPService(container.Redis, 300),
		emailService: appService.NewEmailService(container),
	}
}

func (service *Auth) Login(payload requests.LoginRequest, deviceInfo string, ip string) responses.ServiceResponse {
	var user models.User

	// Check user
	err := service.userRepo.GetUserByUsernameOrEmail(payload.UsernameEmail, payload.UsernameEmail, &user)
	if err == gorm.ErrRecordNotFound {
		return responses.ServiceResponse{
			StatusCode: common.StatusNotFound,
			Message:    "Username/Email not found!",
			Error:      err,
		}
	}
	if err != nil {
		service.container.Logger.Error("auth.service.login-getUser", err.Error(), payload)
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
	authRepo := repository.NewAuth(service.container)
	authHeler := auth.New(service.container, authRepo)
	accessToken, refreshToken, err := authHeler.GenerateToken(user.ID, commonType.User, deviceInfo, ip)
	if err != nil {
		service.container.Logger.Error("auth.service.login-generateToken", err.Error(), user.ID, commonType.User, ip)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      err,
		}
	}

	return responses.ServiceResponse{
		StatusCode: common.StatusSuccess,
		Message:    "Login successful!",
		Data: responses.AuthResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}
}

func (service *Auth) Register(payload requests.RegisterRequest, deviceInfo string, ip string) responses.ServiceResponse {
	var (
		err      error
		isExists bool
		password string
	)

	// Check username
	isExists, err = service.userRepo.UsernameExists(payload.Username)
	if err != nil {
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
			Error:      errors.New("username already exists"),
		}
	}

	// Check email
	isExists, err = service.userRepo.EmailExists(payload.Email)
	if err != nil {
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
			Error:      errors.New("email already exists"),
		}
	}

	avatarRepo := repository.NewAvatar(service.container)
	if err := avatarRepo.GetAvatarByType(payload.AvatarId, commonType.UserAvatar, &models.Avatar{}); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return responses.ServiceResponse{
				StatusCode: common.StatusValidationError,
				Message:    "Avatar not found!",
				Error:      err,
			}
		}

		// Other database errors
		service.container.Logger.Error("auth.service.add-getAvatarByType", err.Error(), payload)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later!",
			Error:      err,
		}
	}

	// Verify OTP
	otpService := appService.NewOTPService(service.container.Redis, 300)
	err = otpService.VerifyOTP(payload.Email, commonType.Register, payload.OTP)
	if err != nil {
		return responses.ServiceResponse{
			StatusCode: common.StatusValidationError,
			Message:    "Invalid OTP!",
			Error:      errors.New("invalid OTP"),
		}
	}

	// Hash password
	password, err = Validate.HashPassword(payload.Password)
	if err != nil {
		service.container.Logger.Error("auth.service.register-hashPassword", err.Error(), payload.Email, payload.Password)
		return responses.ServiceResponse{
			StatusCode: common.StatusValidationError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      err,
		}
	}

	// Create user
	userId, err := service.userRepo.CreateUser(&models.User{
		Name:     payload.Name,
		Username: payload.Username,
		Email:    payload.Email,
		Password: password,
	})
	if err != nil {
		service.container.Logger.Error("auth.service.register-createUser", err.Error(), payload, password)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      err,
		}
	}

	// Create Token
	authRepo := repository.NewAuth(service.container)
	authHeler := auth.New(service.container, authRepo)
	accessToken, refreshToken, err := authHeler.GenerateToken(userId, commonType.User, deviceInfo, ip)
	if err != nil {
		service.container.Logger.Error("auth.service.register-generateToken", err.Error(), userId, commonType.User, ip)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      err,
		}
	}

	// Response
	service.container.Logger.Info("auth.service.register", "User registered successfully!", userId, commonType.User, ip)
	return responses.ServiceResponse{
		StatusCode: common.StatusSuccess,
		Message:    "User registered successfully!",
		Data: responses.AuthResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}
}

func (service *Auth) SentOTP(payload requests.SentOTPRequest) responses.ServiceResponse {

	// Check email
	if payload.Type != commonType.Register {
		isExists, err := service.userRepo.EmailExists(payload.Email)
		if err != nil {
			return responses.ServiceResponse{
				StatusCode: common.StatusServerError,
				Message:    "Oops! Something went wrong. Please try again later.",
				Error:      err,
			}
		}
		if !isExists {
			return responses.ServiceResponse{
				StatusCode: common.StatusValidationError,
				Message:    "Email is not exists.",
				Error:      errors.New("user not found"),
			}
		}
	}

	otp := service.otpService.GenerateOTP()
	if err := service.otpService.SaveOTP(payload.Email, payload.Type, otp); err != nil {
		service.container.Logger.Error("auth-service-sentOtp-redisSaveOTP", err.Error(), payload.Email, payload.Type, otp)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Failed to generate OTP!",
			Error:      err,
		}
	}

	// Send OTP to email
	data := map[string]string{
		"OTP":   otp,
		"Email": payload.Email,
	}
	err := service.emailService.SendEmail(payload.Email, templates.OTP_VERIFICATION_TITLE, templates.OTP_VERIFICATION_TITLE_TEMPLATE, data, []string{})
	if err != nil {
		service.container.Logger.Error("auth.service.sentOtp-emailSend", err.Error(), payload.Email, templates.OTP_VERIFICATION_TITLE, templates.OTP_VERIFICATION_TITLE_TEMPLATE, data)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Failed to send OTP!",
			Error:      err,
		}
	}

	service.container.Logger.Info("auth.service.sentOtp", "OTP sent successfully to the email address!", payload.Email, templates.OTP_VERIFICATION_TITLE, templates.OTP_VERIFICATION_TITLE_TEMPLATE, data)
	return responses.ServiceResponse{
		StatusCode: common.StatusSuccess,
		Message:    "OTP sent successfully to the email address!",
	}
}

func (service *Auth) ResetPassword(payload requests.ResetPasswordRequest, deviceInfo string, ip string) responses.ServiceResponse {
	var user models.User

	// Check user
	err := service.userRepo.GetUserByUsernameOrEmail("", payload.Email, &user)
	if err == gorm.ErrRecordNotFound {
		return responses.ServiceResponse{
			StatusCode: common.StatusNotFound,
			Message:    "User not found!",
			Error:      err,
		}
	}
	if err != nil {
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      err,
		}
	}

	// Validate old and password new password
	if err := Validate.VerifyPassword(user.Password, payload.Password); err == nil {
		return responses.ServiceResponse{
			StatusCode: common.StatusValidationError,
			Message:    "Password is the same as the previous one!",
			Error:      errors.New("auth issue"),
		}
	}

	// Verify OTP
	err = service.otpService.VerifyOTP(payload.Email, commonType.ResetPassword, payload.OTP)
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
		service.container.Logger.Error("auth.service.resetPassword-hashPassword", err.Error(), payload)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      err,
		}
	}

	// Update password in user table
	err = service.userRepo.UpdatePasswordByEmail(payload.Email, password)
	if err != nil {
		service.container.Logger.Error("auth.service.resetPassword-updatePasswordByEmail", err.Error(), payload.Email, password)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    err.Error(),
			Error:      err,
		}
	}

	// Create Token
	authRepo := repository.NewAuth(service.container)
	authHeler := auth.New(service.container, authRepo)
	accessToken, refreshToken, err := authHeler.GenerateToken(user.ID, commonType.User, deviceInfo, ip)
	if err != nil {
		service.container.Logger.Error("auth.service.resetPassword-generateToken", err.Error(), user.ID, commonType.User, ip)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      err,
		}
	}

	return responses.ServiceResponse{
		StatusCode: common.StatusSuccess,
		Message:    "Password updated successfully!",
		Data: responses.AuthResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}
}

func (service *Auth) GetToken(payload requests.TokenRequest, deviceInfo string, ip string) responses.ServiceResponse {
	var user models.User

	authRepo := repository.NewAuth(service.container)
	authHeler := auth.New(service.container, authRepo)
	tokenClaims, err := authHeler.VerifyRefreshToken(payload.RefreshToken)
	if err != nil {
		return responses.ServiceResponse{
			StatusCode: common.StatusValidationError,
			Message:    err.Error(),
			Error:      err,
		}
	}

	// Check user
	claims, _ := tokenClaims.(*auth.AuthClaim)
	err = service.userRepo.GetUser(claims.SessionID, &user)
	if err == gorm.ErrRecordNotFound {
		return responses.ServiceResponse{
			StatusCode: common.StatusNotFound,
			Message:    "User not found!",
			Error:      err,
		}
	}
	if err != nil {
		service.container.Logger.Error("auth.service.getToken-getUser", err.Error(), payload, claims)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      err,
		}
	}

	// Create Token
	accessToken, refreshToken, err := authHeler.GenerateToken(user.ID, commonType.User, deviceInfo, ip)
	if err != nil {
		service.container.Logger.Error("auth.service.getToken-generateToken", err.Error(), user.ID, commonType.User, ip)
		return responses.ServiceResponse{
			StatusCode: common.StatusServerError,
			Message:    "Oops! Something went wrong. Please try again later.",
			Error:      err,
		}
	}

	// Response
	return responses.ServiceResponse{
		StatusCode: common.StatusSuccess,
		Message:    "Token re-generated successfully!",
		Data: responses.AuthResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}
}

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
	userRepo     repository.UserRepository
	authRepo     repository.AuthRepository
	avatarRepo   repository.AvatarRepository
	authHelper   *auth.Auth
	otpService   appService.OTPService
	emailService appService.EmailService
}

func NewAuth(container *storage.Container) *Auth {
	authRepo := repository.NewAuth(container)
	return &Auth{
		container:    container,
		userRepo:     repository.NewUser(container),
		authRepo:     authRepo,
		avatarRepo:   repository.NewAvatar(container),
		authHelper:   auth.New(container, authRepo),
		otpService:   appService.NewOTP(container.Redis, 300),
		emailService: appService.NewEmail(container),
	}
}

func (service *Auth) Login(rctx *requests.RequestContext, payload requests.LoginRequest, deviceInfo string, ip string) responses.ServiceResponse {
	var user models.User

	// Check user
	if err := service.userRepo.GetUserByUsernameOrEmail(rctx, payload.UsernameEmail, payload.UsernameEmail, &user); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "Username/Email not found!", err)
		}

		service.container.Logger.Error("auth.service.login-GetUserByUsernameOrEmail", "error", err.Error(), "payload", payload)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong. Please try again later", err)
	}

	// Validate password
	if err := Validate.VerifyPassword(user.Password, payload.Password); err != nil {
		return responses.ErrorResponse(common.StatusValidationError, "Invalid password!", err)
	}

	// Create Token
	accessToken, refreshToken, err := service.authHelper.GenerateToken(rctx, user.ID, commonType.UserTypeUser, deviceInfo, ip)
	if err != nil {
		service.container.Logger.Error("auth.service.login-GenerateToken", "error", err.Error(), "userId", user.ID, "userType", commonType.UserTypeUser, "deviceInfo", deviceInfo, "ip", ip)
		return responses.ErrorResponse(common.StatusServerError, "Failed to generate tokens. Please try again later.", err)
	}

	// Response
	return responses.SuccessResponse("Login successful!", responses.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func (service *Auth) Register(rctx *requests.RequestContext, payload requests.RegisterRequest, deviceInfo string, ip string) responses.ServiceResponse {
	var password string

	// Check username
	if err := service.userRepo.UsernameExists(rctx, payload.Username); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			service.container.Logger.Error("auth.service.register-UsernameExists", "error", err.Error(), "username", payload.Username)
			return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong. Please try again later.", err)
		}
	} else {
		return responses.ErrorResponse(common.StatusValidationError, "Username already exists!", errors.New("username already exists"))
	}

	// Check email
	if err := service.userRepo.EmailExists(rctx, payload.Email); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			service.container.Logger.Error("auth.service.register-EmailExists", "error", err.Error(), "email", payload.Email)
			return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong. Please try again later.", err)
		}
	} else {
		return responses.ErrorResponse(common.StatusValidationError, "Email already exists!", errors.New("email already exists"))
	}

	// Verify avatar
	if err := service.avatarRepo.AvatarByTypeExists(rctx, payload.AvatarId, commonType.AvatarTypeUser); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusValidationError, "Avatar not found!.", errors.New("avatar not found"))
		}
		service.container.Logger.Error("auth.service.register-AvatarByTypeExists", "error", err.Error(), "avatarId", payload.AvatarId, "avatarType", commonType.AvatarTypeUser)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong. Please try again later.", err)
	}

	// Verify OTP
	if err := service.otpService.VerifyOTP(payload.Email, commonType.OtpTypeRegister, payload.OTP); err != nil {
		return responses.ErrorResponse(common.StatusValidationError, "Invalid OTP!", errors.New("invalid otp"))
	}

	// Hash password
	password, err := Validate.HashPassword(payload.Password)
	if err != nil {
		service.container.Logger.Error("auth.service.register-HashPassword", "error", err.Error(), "email", payload.Email, "password", payload.Password)
		return responses.ErrorResponse(common.StatusServerError, "Oops! Something went wrong. Please try again later.", err)
	}

	// Create user
	userId, err := service.userRepo.CreateUser(rctx, &models.User{
		Name:     payload.Name,
		Username: payload.Username,
		Email:    payload.Email,
		Password: password,
	})
	if err != nil {
		service.container.Logger.Error("auth.service.register-CreateUser", "error", err.Error(), "payload", payload)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong. Please try again later.", err)
	}

	// Create Token
	var accessToken, refreshToken string
	if accessToken, refreshToken, err = service.authHelper.GenerateToken(rctx, userId, commonType.UserTypeUser, deviceInfo, ip); err != nil {
		service.container.Logger.Error("auth.service.register-GenerateToken", "error", err.Error(), "userId", userId, "type", commonType.UserTypeUser)
		return responses.ErrorResponse(common.StatusServerError, "Failed to generate tokens. Please try again later.", err)
	}

	// Response
	service.container.Logger.Info("auth.service.register.success", "messgae", "User registered successfully!", "userId", userId, "type", commonType.UserTypeUser, "ip", ip)
	return responses.SuccessResponse("Register successfully!", responses.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func (service *Auth) SendOTP(rctx *requests.RequestContext, payload requests.SentOTPRequest) responses.ServiceResponse {

	// Check email
	if payload.Type != commonType.OtpTypeRegister {
		if err := service.userRepo.EmailExists(rctx, payload.Email); err != nil {
			if err == gorm.ErrRecordNotFound {
				return responses.ErrorResponse(common.StatusValidationError, "Email is not exists!", errors.New("email is not exists"))
			}
			service.container.Logger.Error("auth.service.sendOTP-EmailExists", "error", err.Error(), "email", payload.Email)
			return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong. Please try again later.", err)
		}
	}

	// OTP generate and save
	otp := service.otpService.GenerateOTP()
	if err := service.otpService.SaveOTP(payload.Email, payload.Type, otp); err != nil {
		service.container.Logger.Error("auth.service.sendOTP-SaveOTP", "error", err.Error(), "email", payload.Email, "type", payload.Type)
		return responses.ErrorResponse(common.StatusServerError, "Failed to generate OTP.", err)
	}

	// Send OTP to email
	data := map[string]string{
		"OTP":   otp,
		"Email": payload.Email,
	}
	if err := service.emailService.SendEmail(payload.Email, templates.OTP_VERIFICATION_TITLE, templates.OTP_VERIFICATION_TITLE_TEMPLATE, data, []string{}); err != nil {
		service.container.Logger.Error("auth.service.sendOTP-SendEmail", "error", err.Error(), "email", payload.Email, "templateName", templates.OTP_VERIFICATION_TITLE, "templatePath", templates.OTP_VERIFICATION_TITLE_TEMPLATE, "data", data)
		return responses.ErrorResponse(common.StatusServerError, "Failed to send OTP!", err)
	}

	// Response
	service.container.Logger.Info("auth.service.sendOTP.success", "messgae", "OTP sent successfully to the email address!", "email", payload.Email, "type", payload.Type)
	return responses.SuccessResponse("OTP sent successfully to the email address!", nil)
}

func (service *Auth) ResetPassword(rctx *requests.RequestContext, payload requests.ResetPasswordRequest, deviceInfo string, ip string) responses.ServiceResponse {
	var user models.User

	// Check user
	if err := service.userRepo.GetUserByUsernameOrEmail(rctx, "", payload.Email, &user); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "User not found!", err)
		}
		service.container.Logger.Error("auth.service.resetPassword-GetUserByUsernameOrEmail", "error", err.Error(), "email", payload.Email)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong. Please try again later.", err)
	}

	// Validate old and password new password
	if err := Validate.VerifyPassword(user.Password, payload.Password); err == nil {
		return responses.ErrorResponse(common.StatusValidationError, "Password is the same as the previous one!", errors.New("password is the same as the previous one"))
	}

	// Verify OTP
	if err := service.otpService.VerifyOTP(payload.Email, commonType.OtpTypeResetPassword, payload.OTP); err != nil {
		return responses.ErrorResponse(common.StatusValidationError, "Invalid OTP!", err)
	}

	// Hash password
	password, err := Validate.HashPassword(payload.Password)
	if err != nil {
		service.container.Logger.Error("auth.service.resetPassword-HashPassword", "error", err.Error(), "email", payload.Email, "password", payload.Password)
		return responses.ErrorResponse(common.StatusServerError, "Oops! Something went wrong. Please try again later.", err)
	}

	// Update password
	if err := service.userRepo.UpdatePasswordByEmail(rctx, payload.Email, password); err != nil {
		service.container.Logger.Error("auth.service.resetPassword-UpdatePasswordByEmail", "error", err.Error(), "email", payload.Email, "password", payload.Password)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong. Please try again later.", err)
	}

	// Create Token
	var accessToken, refreshToken string
	if accessToken, refreshToken, err = service.authHelper.GenerateToken(rctx, user.ID, commonType.UserTypeUser, deviceInfo, ip); err != nil {
		service.container.Logger.Error("auth.service.resetPassword-UpdatePasswordByEmail", "error", err.Error(), "userId", user.ID, "password", "userType", commonType.UserTypeUser)
		return responses.ErrorResponse(common.StatusServerError, "Failed to generate tokens. Please try again later.", err)
	}

	// Response
	return responses.SuccessResponse("Password updated successfully!", responses.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func (service *Auth) GetToken(rctx *requests.RequestContext, payload requests.TokenRequest, deviceInfo string, ip string) responses.ServiceResponse {
	var user models.User

	// Verify refreshToken
	tokenClaims, err := service.authHelper.VerifyRefreshToken(rctx, payload.RefreshToken)
	if err != nil {
		return responses.ErrorResponse(common.StatusValidationError, err.Error(), err)
	}

	// Check user
	claims, _ := tokenClaims.(*auth.AuthClaim)
	if err = service.userRepo.GetUser(rctx, claims.UserId, &user); err != nil {
		if err == gorm.ErrRecordNotFound {
			return responses.ErrorResponse(common.StatusNotFound, "User not found!", err)
		}

		service.container.Logger.Error("auth.service.getToken-GetUser", "error", err.Error(), "userId", claims.UserId)
		return responses.ErrorResponse(common.StatusDatabaseError, "Oops! Something went wrong. Please try again later.", err)
	}

	// Remove session
	service.authRepo.DeleteSession(rctx, claims.SessionID)

	// Create Token
	var accessToken, refreshToken string
	if accessToken, refreshToken, err = service.authHelper.GenerateToken(rctx, user.ID, commonType.UserTypeUser, deviceInfo, ip); err != nil {
		service.container.Logger.Error("auth.service.getToken-GenerateToken", "error", err.Error(), "userId", user.ID, "userType", commonType.UserTypeUser, "deviceInfo", deviceInfo, "ip", ip)
		return responses.ErrorResponse(common.StatusServerError, "Failed to generate tokens. Please try again later.", err)
	}

	// Response
	return responses.SuccessResponse("Token re-generated successfully!", responses.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

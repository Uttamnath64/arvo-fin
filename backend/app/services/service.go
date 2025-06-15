package services

import (
	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/responses"
	"github.com/Uttamnath64/arvo-fin/pkg/validater"
)

var (
	Validate *validater.Validater
)

type EmailService interface {
	SendEmail(to, subject, templateFile string, data map[string]string, attachments []string) error
}

type OTPService interface {
	GenerateOTP() string
	SaveOTP(email string, otpType commonType.OtpType, otp string) error
	VerifyOTP(email string, otpType commonType.OtpType, providedOTP string) error
}

type PortfolioService interface {
	GetList(userId uint, userType commonType.UserType) responses.ServiceResponse
	Get(id, userId uint, userType commonType.UserType) responses.ServiceResponse
}

type UserService interface {
	Get(userId uint) responses.ServiceResponse
	GetSettings(userId uint) responses.ServiceResponse
	Update(payload requests.MeRequest, userId uint) responses.ServiceResponse
	UpdateSettings(payload requests.SettingsRequest, userId uint) responses.ServiceResponse
}

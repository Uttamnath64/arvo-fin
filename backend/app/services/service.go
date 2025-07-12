package services

import (
	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/responses"
	"github.com/Uttamnath64/arvo-fin/pkg/pagination"
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
	GetList(rctx *requests.RequestContext, userId uint) responses.ServiceResponse
	Get(rctx *requests.RequestContext, id, userId uint, userType commonType.UserType) responses.ServiceResponse
}

type UserService interface {
	Get(rctx *requests.RequestContext, userId uint) responses.ServiceResponse
	GetSettings(rctx *requests.RequestContext, userId uint) responses.ServiceResponse
	Update(rctx *requests.RequestContext, payload requests.MeRequest, userId uint) responses.ServiceResponse
	UpdateSettings(rctx *requests.RequestContext, payload requests.SettingsRequest, userId uint) responses.ServiceResponse
}

type AvatarService interface {
	Get(rctx *requests.RequestContext, id uint) responses.ServiceResponse
	GetAvatarsByType(rctx *requests.RequestContext, avatarType commonType.AvatarType) responses.ServiceResponse
	Create(rctx *requests.RequestContext, payload requests.AvatarRequest) responses.ServiceResponse
	Update(rctx *requests.RequestContext, id uint, payload requests.AvatarRequest) responses.ServiceResponse
}

type AccountService interface {
	Get(rctx *requests.RequestContext, id uint) responses.ServiceResponse
	GetList(rctx *requests.RequestContext, portfolioId, userId uint) responses.ServiceResponse
	AccountTypes(rctx *requests.RequestContext) responses.ServiceResponse
	Create(rctx *requests.RequestContext, userId uint, payload requests.AccountRequest) responses.ServiceResponse
	Update(rctx *requests.RequestContext, id, userId uint, payload requests.AccountUpdateRequest) responses.ServiceResponse
	Delete(rctx *requests.RequestContext, id, userId uint) responses.ServiceResponse
}

type CategoryService interface {
	Get(rctx *requests.RequestContext, id uint) responses.ServiceResponse
	GetList(rctx *requests.RequestContext, portfolioId, userId uint) responses.ServiceResponse
	Create(rctx *requests.RequestContext, payload requests.CategoryRequest) responses.ServiceResponse
	Update(rctx *requests.RequestContext, id uint, payload requests.CategoryRequest) responses.ServiceResponse
	Delete(rctx *requests.RequestContext, portfolioId, id uint) responses.ServiceResponse
}

type TransactionService interface {
	Get(rctx *requests.RequestContext, id uint) responses.ServiceResponse
	GetList(rctx *requests.RequestContext, transactionQuery requests.TransactionQuery, pagination pagination.Pagination) responses.ServiceResponse
	Create(rctx *requests.RequestContext, payload requests.TransactionRequest) responses.ServiceResponse
	Update(rctx *requests.RequestContext, id uint, payload requests.TransactionRequest) responses.ServiceResponse
	Delete(rctx *requests.RequestContext, id uint) responses.ServiceResponse
}

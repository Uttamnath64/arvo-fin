package repository

import (
	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/responses"
	"github.com/Uttamnath64/arvo-fin/pkg/pagination"
)

type AuthRepository interface {
	GetSessionByUser(rctx *requests.RequestContext, userId uint, userType commonType.UserType, signedToken string) (*models.Session, error)
	GetSessionByRefreshToken(rctx *requests.RequestContext, refreshToken string, userType commonType.UserType) (*models.Session, error)
	CreateSession(rctx *requests.RequestContext, session *models.Session) (uint, error)
	UpdateSession(rctx *requests.RequestContext, id uint, refreshToken string, expiresAt int64) error
	DeleteSession(rctx *requests.RequestContext, sessionID uint) error
}

type AvatarRepository interface {
	Get(rctx *requests.RequestContext, id uint) (*models.Avatar, error)
	GetByNameAndType(rctx *requests.RequestContext, name string, avatarType commonType.AvatarType) *models.Avatar
	AvatarByTypeExists(rctx *requests.RequestContext, id uint, avatarType commonType.AvatarType) error
	GetAvatarsByType(rctx *requests.RequestContext, avatarType commonType.AvatarType) (*[]models.Avatar, error)
	Create(rctx *requests.RequestContext, payload models.Avatar) (uint, error)
	Update(rctx *requests.RequestContext, id uint, payload requests.AvatarRequest) error
}

type PortfolioRepository interface {
	UserPortfolioExists(rctx *requests.RequestContext, id, userId uint) error
	GetList(rctx *requests.RequestContext, userId uint) (*[]models.Portfolio, error)
	Get(rctx *requests.RequestContext, id, userId uint, userType commonType.UserType) (*models.Portfolio, error)
	Create(rctx *requests.RequestContext, portfolio models.Portfolio) error
	Update(rctx *requests.RequestContext, id, userId uint, payload requests.PortfolioRequest) error
	Delete(rctx *requests.RequestContext, id, userId uint) error
}

type UserRepository interface {
	GetUserByUsernameOrEmail(rctx *requests.RequestContext, username string, email string, user *models.User) error
	UsernameExists(rctx *requests.RequestContext, username string) error
	EmailExists(rctx *requests.RequestContext, email string) error
	CreateUser(rctx *requests.RequestContext, user *models.User) (uint, error)
	UpdatePasswordByEmail(rctx *requests.RequestContext, email, newPassword string) error
	GetUser(rctx *requests.RequestContext, userId uint, user *models.User) error
	Get(rctx *requests.RequestContext, userId uint) (*responses.MeResponse, error)
	GetSettings(rctx *requests.RequestContext, userId uint) (*responses.SettingsResponse, error)
	Update(rctx *requests.RequestContext, userId uint, payload requests.MeRequest) error
	UpdateSettings(rctx *requests.RequestContext, userId uint, payload requests.SettingsRequest) error
}

type CurrencyRepository interface {
	CodeExists(rctx *requests.RequestContext, code string) error
}

type AccountRepository interface {
	UserAccountExists(rctx *requests.RequestContext, id, portfolioId, userId uint) error
	GetList(rctx *requests.RequestContext, portfolioId, userId uint) (*[]models.Account, error)
	Get(rctx *requests.RequestContext, id uint) (*models.Account, error)
	Create(rctx *requests.RequestContext, account models.Account) (uint, error)
	Update(rctx *requests.RequestContext, id, userId uint, payload requests.AccountUpdateRequest) error
	Delete(rctx *requests.RequestContext, id, userId uint) error
}

type CategoryRepository interface {
	UserCategoryExists(rctx *requests.RequestContext, id, portfolioId, userId uint) error
	GetList(rctx *requests.RequestContext, portfolioId, userId uint) (*[]models.Category, error)
	Get(rctx *requests.RequestContext, id, userId uint) (*models.Category, error)
	Create(rctx *requests.RequestContext, category models.Category) (uint, error)
	Update(rctx *requests.RequestContext, id, userId uint, payload requests.CategoryRequest) error
	Delete(rctx *requests.RequestContext, id, userId uint) error
}

type TransactionRepository interface {
	Get(rctx *requests.RequestContext, id uint) (*models.Transaction, error)
	GetList(rctx *requests.RequestContext, transactionQuery requests.TransactionQuery, pagination pagination.Pagination) (*[]models.Transaction, error)
	Create(rctx *requests.RequestContext, transaction models.Transaction) (uint, error)
	Update(rctx *requests.RequestContext, id uint, payload requests.TransactionRequest) error
	Delete(rctx *requests.RequestContext, id uint) error
}

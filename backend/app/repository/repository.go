package repository

import (
	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/responses"
)

type AuthRepository interface {
	GetSessionByUser(userId uint, userType commonType.UserType, signedToken string) (*models.Session, error)
	GetSessionByRefreshToken(refreshToken string, userType commonType.UserType) (*models.Session, error)
	CreateSession(session *models.Session) error
	UpdateSession(id uint, refreshToken string, expiresAt int64) error
	DeleteSession(sessionID uint) error
}

type AvatarRepository interface {
	GetAvatar(id uint, avatar *models.Avatar) error
	GetAvatarByType(id uint, avatarType commonType.AvatarType, avatar *models.Avatar) error
}

type PortfolioRepository interface {
	GetList(userId uint, userType commonType.UserType) (*[]responses.PortfolioResponse, error)
	Get(id, userId uint, userType commonType.UserType) (*responses.PortfolioResponse, error)
	Add(portfolio models.Portfolio) error
	Update(id, userId uint, payload requests.PortfolioRequest) error
	Delete(id, userId uint) error
}

type UserRepository interface {
	GetUserByUsernameOrEmail(username string, email string, user *models.User) error
	UsernameExists(username string) (bool, error)
	EmailExists(email string) (bool, error)
	CreateUser(user *models.User) (uint, error)
	UpdatePasswordByEmail(email, newPassword string) error
	GetUser(userId uint, user *models.User) error
}

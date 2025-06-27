package auth

import (
	"errors"
	"fmt"
	"time"

	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/repository"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"github.com/golang-jwt/jwt"
)

type Auth struct {
	container *storage.Container
	authRepo  repository.AuthRepository
}

type AuthClaim struct {
	UserId    uint                `json:"user_id"`
	UserType  commonType.UserType `json:"user_type"`
	SessionID uint                `json:"session_id"`
	jwt.StandardClaims
}

func New(container *storage.Container, authRepo repository.AuthRepository) *Auth {
	return &Auth{
		container: container,
		authRepo:  authRepo,
	}
}

func (auth *Auth) GenerateToken(rctx *requests.RequestContext, userId uint, userType commonType.UserType, deviceInfo, ipAddress string) (string, string, error) {

	var accessExpiresAt = time.Now().Add(auth.container.Env.Auth.AccessTokenExpired).Unix()
	var refreshExpiresAt = time.Now().Add(auth.container.Env.Auth.RefreshTokenExpired).Unix()

	// create settion
	sessionId, err := auth.authRepo.CreateSession(rctx, &models.Session{
		UserID:     userId,
		UserType:   userType,
		DeviceInfo: deviceInfo,
		IPAddress:  ipAddress,
	})
	if err != nil {
		return "", "", err
	}

	accessTokenJWT := jwt.NewWithClaims(jwt.SigningMethodRS256, &AuthClaim{
		UserId:    userId,
		UserType:  userType,
		SessionID: sessionId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessExpiresAt,
		},
	})

	accessToken, err := accessTokenJWT.SignedString(auth.container.Env.Auth.AccessPrivateKey)
	if err != nil {
		return "", "", err
	}

	refreshTokenJWT := jwt.NewWithClaims(jwt.SigningMethodRS256, &AuthClaim{
		UserId:    userId,
		UserType:  userType,
		SessionID: sessionId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshExpiresAt,
		},
	})

	refreshToken, err := refreshTokenJWT.SignedString(auth.container.Env.Auth.RefreshPrivateKey)
	if err != nil {
		return "", "", err
	}

	if err := auth.authRepo.UpdateSession(rctx, sessionId, refreshToken, refreshExpiresAt); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (auth *Auth) VerifyRefreshToken(rctx *requests.RequestContext, refreshToken string) (interface{}, error) {

	token, err := jwt.ParseWithClaims(
		refreshToken,
		&AuthClaim{},
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return auth.container.Env.Auth.RefreshPublicKey, nil
		},
	)
	if err != nil || !token.Valid {
		return nil, errors.New("refresh token is invalid")
	}

	claims, ok := token.Claims.(*AuthClaim)
	if !ok || claims.SessionID == 0 {
		return nil, errors.New("invalid refresh token claims")
	}

	if err := auth.isValidRefreshToken(rctx, claims.SessionID, claims.UserType, refreshToken); err != nil {
		return nil, err
	}

	return claims, nil
}

func (auth *Auth) isValidRefreshToken(rctx *requests.RequestContext, sessionID uint, userType commonType.UserType, refreshToken string) error {
	session, err := auth.authRepo.GetSessionByRefreshToken(rctx, refreshToken, userType)
	if err != nil {
		return err
	}

	// Check if token exists
	if session == nil {
		return errors.New("refresh token not found")
	}

	if session.ExpiresAt < time.Now().Unix() {
		return errors.New("refresh token is expired")
	}

	return nil
}

package auth

import (
	"encoding/base64"
	"errors"
	"time"

	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/repository"
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

func (auth *Auth) GenerateToken(userId uint, userType commonType.UserType, deviceInfo, ipAddress string) (string, string, error) {

	var accessExpiresAt = time.Now().Add(auth.container.Env.Auth.AccessTokenExpired * time.Hour).Unix()
	var refreshExpiresAt = time.Now().Add(auth.container.Env.Auth.RefreshTokenExpired * time.Hour).Unix()

	// AccessPrivateKey
	decodedAccessPrivateKey, err := base64.StdEncoding.DecodeString(auth.container.Env.Auth.AccessTokenPrivateKey)
	if err != nil {
		return "", "", errors.New("Could not decode key: " + err.Error())
	}
	AccessPrivateKey, err := jwt.ParseRSAPrivateKeyFromPEM(decodedAccessPrivateKey)
	if err != nil {
		return "", "", errors.New("Could not parse key: " + err.Error())
	}

	// RefreshPrivateKey
	decodedRefreshPrivateKey, err := base64.StdEncoding.DecodeString(auth.container.Env.Auth.RefreshTokenPrivateKey)
	if err != nil {
		return "", "", errors.New("Could not decode key: " + err.Error())
	}
	RefreshPrivateKey, err := jwt.ParseRSAPrivateKeyFromPEM(decodedRefreshPrivateKey)

	if err != nil {
		return "", "", errors.New("Could not parse key: " + err.Error())
	}

	// create settion
	session := models.Session{
		UserID:     userId,
		UserType:   userType,
		DeviceInfo: deviceInfo,
		IPAddress:  ipAddress,
	}
	if err := auth.authRepo.CreateSession(&session); err != nil {
		return "", "", err
	}

	accessTokenJWT := jwt.NewWithClaims(jwt.SigningMethodRS256, &AuthClaim{
		UserId:    userId,
		UserType:  userType,
		SessionID: session.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessExpiresAt,
		},
	})

	accessToken, err := accessTokenJWT.SignedString(AccessPrivateKey)
	if err != nil {
		return "", "", err
	}

	refreshTokenJWT := jwt.NewWithClaims(jwt.SigningMethodRS256, &AuthClaim{
		UserId:    userId,
		UserType:  userType,
		SessionID: session.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshExpiresAt,
		},
	})

	refreshToken, err := refreshTokenJWT.SignedString(RefreshPrivateKey)
	if err != nil {
		return "", "", err
	}

	if err := auth.authRepo.UpdateSession(session.ID, refreshToken, refreshExpiresAt); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (auth *Auth) VerifyRefreshToken(refreshToken string) (interface{}, error) {

	decodedRefreshPublicKey, err := base64.StdEncoding.DecodeString(auth.container.Env.Auth.RefreshTokenPublicKey)
	if err != nil {
		auth.container.Logger.Error("auth.service.GetToken-VerifyRefreshToken", err.Error())
		return nil, errors.New("Refresh token is invalid!")
	}

	RefreshPublicKey, err := jwt.ParseRSAPublicKeyFromPEM(decodedRefreshPublicKey)
	if err != nil {
		auth.container.Logger.Error("auth.service.GetToken-VerifyRefreshToken", err.Error())
		return "", errors.New("Refresh token is invalid!")
	}

	token, err := jwt.ParseWithClaims(
		refreshToken,
		&AuthClaim{},
		func(t *jwt.Token) (interface{}, error) {
			return RefreshPublicKey, nil
		},
	)
	if err != nil {
		return nil, errors.New("Refresh token is invalid!")
	}

	claims, ok := token.Claims.(*AuthClaim)
	if !ok {
		return nil, errors.New("Refresh token is invalid!")
	}

	if err := auth.isValidRefreshToken(claims.SessionID, claims.UserType, refreshToken); err != nil {
		return nil, err
	}

	return claims, nil
}

func (auth *Auth) isValidRefreshToken(sessionID uint, userType commonType.UserType, refreshToken string) error {
	session, err := auth.authRepo.GetSessionByRefreshToken(refreshToken, userType)
	if err != nil {
		return err
	}

	// Check if token exists
	if session == nil {
		return errors.New("Refresh token not found!")
	}

	if session.ExpiresAt < time.Now().Unix() {
		return errors.New("Refresh token is expired!")
	}

	return nil
}

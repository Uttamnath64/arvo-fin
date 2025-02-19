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
	authRepo  *repository.Auth
}

type AuthClaim struct {
	ReferenceId uint                `json:"referenceId"`
	UserType    commonType.UserType `json:"userType"`
	jwt.StandardClaims
}

func New(container *storage.Container, authRepo *repository.Auth) *Auth {
	return &Auth{
		container: container,
		authRepo:  repository.NewAuth(container),
	}
}

func (auth *Auth) GenerateToken(referenceId uint, userType commonType.UserType, ip string) (string, string, error) {

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

	accessTokenJWT := jwt.NewWithClaims(jwt.SigningMethodRS256, &AuthClaim{
		ReferenceId: referenceId,
		UserType:    userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessExpiresAt,
		},
	})

	accessToken, err := accessTokenJWT.SignedString(AccessPrivateKey)
	if err != nil {
		return "", "", err
	}

	refreshTokenJWT := jwt.NewWithClaims(jwt.SigningMethodRS256, &AuthClaim{
		ReferenceId: referenceId,
		UserType:    userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshExpiresAt,
		},
	})

	refreshToken, err := refreshTokenJWT.SignedString(RefreshPrivateKey)
	if err != nil {
		return "", "", err
	}

	if err := auth.CreateToken(&models.Token{
		ReferenceId: referenceId,
		UserType:    userType,
		IP:          ip,
		Token:       refreshToken,
		ExpiresAt:   refreshExpiresAt,
	}); err != nil {
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

	if err := auth.isValidRefreshToken(claims.ReferenceId, claims.UserType, refreshToken); err != nil {
		return nil, err
	}

	return claims, nil
}

func (auth *Auth) isValidRefreshToken(referenceID uint, userType commonType.UserType, refreshToken string) error {
	token, err := auth.authRepo.GetTokenByReference(referenceID, userType, refreshToken)
	if err != nil {
		return err
	}

	// Check if token exists
	if token == nil {
		return errors.New("Refresh token not found!")
	}

	if token.ExpiresAt < time.Now().Unix() {
		return errors.New("Refresh token is expired!")
	}

	return nil
}

func (auth *Auth) CreateToken(token *models.Token) error {
	return auth.authRepo.CreateToken(token)
}

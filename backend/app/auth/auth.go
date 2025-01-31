package auth

import (
	"encoding/base64"
	"errors"
	"time"

	"github.com/Uttamnath64/arvo-fin/app/config"
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/services"
	"github.com/Uttamnath64/arvo-fin/pkg/logger"
	"github.com/golang-jwt/jwt"
)

type Auth struct {
	config *config.Config
	env    *config.AppEnv
	logger *logger.Logger
}

type AuthClaim struct {
	ReferenceId uint `json:"referenceId"`
	UserType    byte `json:"userType"`
	jwt.StandardClaims
}

func New(con *config.Config, env *config.AppEnv, logger *logger.Logger) *Auth {
	return &Auth{
		config: con,
		env:    env,
		logger: logger,
	}
}

func (auth *Auth) GenerateToken(referenceId uint, userType byte, ip string) (string, string, error) {

	var accessExpiresAt = time.Now().Add(auth.env.Auth.AccessTokenExpired * time.Hour).Unix()
	var refreshExpiresAt = time.Now().Add(auth.env.Auth.RefreshTokenExpired * time.Hour).Unix()

	// AccessPrivateKey
	decodedAccessPrivateKey, err := base64.StdEncoding.DecodeString(auth.env.Auth.AccessTokenPrivateKey)
	if err != nil {
		return "", "", errors.New("Could not decode key: " + err.Error())
	}
	AccessPrivateKey, err := jwt.ParseRSAPrivateKeyFromPEM(decodedAccessPrivateKey)
	if err != nil {
		return "", "", errors.New("Could not parse key: " + err.Error())
	}

	// RefreshPrivateKey
	decodedRefreshPrivateKey, err := base64.StdEncoding.DecodeString(auth.env.Auth.RefreshTokenPrivateKey)
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

	if err := services.NewAuthService(auth.config, auth.logger).AddToken(&models.Token{
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

func (auth *Auth) VerifyRefreshToken(signedToken string) (interface{}, error) {

	decodedRefreshPublicKey, err := base64.StdEncoding.DecodeString(auth.env.Auth.RefreshTokenPublicKey)
	if err != nil {
		return nil, errors.New("Could not decode: " + err.Error())
	}

	RefreshPublicKey, err := jwt.ParseRSAPublicKeyFromPEM(decodedRefreshPublicKey)

	if err != nil {
		return "", errors.New("Could not parse key: " + err.Error())
	}

	token, err := jwt.ParseWithClaims(
		signedToken,
		&AuthClaim{},
		func(t *jwt.Token) (interface{}, error) {
			return RefreshPublicKey, nil
		},
	)

	if err != nil {
		return nil, errors.New("ParseWithClaims Error: " + err.Error())
	}

	claims, ok := token.Claims.(*AuthClaim)
	if !ok {
		return nil, errors.New("Couldn't parse claims")
	}

	if err := services.NewAuthService(auth.config, auth.logger).IsValidRefreshToken(claims.ReferenceId, claims.UserType, signedToken); err != nil {
		return nil, errors.New("Refresh token is invalid")
	}

	return claims, nil
}

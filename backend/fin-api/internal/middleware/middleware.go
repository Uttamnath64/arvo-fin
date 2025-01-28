package middleware

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/Uttamnath64/arvo-fin/app/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Middleware struct {
	config *config.Config
	env    *config.AppEnv
}

func New(con *config.Config, env *config.AppEnv) *Middleware {
	return &Middleware{
		config: con,
		env:    env,
	}
}

func (auth *Middleware) Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing access token!"})
			ctx.Abort()
			return
		}

		// remove bearer
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		// conver key
		decodedAccessPublicKey, err := base64.StdEncoding.DecodeString(auth.env.Auth.AccessTokenPublicKey)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Could not decode key: " + err.Error()})
		}

		AccessPublicKey, err := jwt.ParseRSAPublicKeyFromPEM(decodedAccessPublicKey)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Could not parse key: " + err.Error()})
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return AccessPublicKey, nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid access token! Error: " + err.Error(), "accessToken": tokenString})
			ctx.Abort()
			return
		}

		// Check if token claims exist and have the expected format
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			// Log the token payload for debugging purposes
			fmt.Println("Invalid token claims format! Token payload:", token.Claims)
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims format!"})
			ctx.Abort()
			return
		}

		referenceID, ok := claims["referenceId"].(float64)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid referenceId format!"})
			ctx.Abort()
			return
		}

		userType, ok := claims["userType"].(float64)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid userType format!"})
			ctx.Abort()
			return
		}

		ctx.Set("ReferenceID", int(referenceID))
		ctx.Set("UserType", int(userType))
		ctx.Next()
	}
}

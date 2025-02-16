package middleware

import (
	"encoding/base64"
	"net/http"
	"strings"

	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/responses"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Middleware struct {
	container *storage.Container
}

func New(container *storage.Container) *Middleware {
	return &Middleware{
		container: container,
	}
}

func (m *Middleware) Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, responses.ApiResponse{
				Status:  false,
				Message: "Missing access token!",
			})
			ctx.Abort()
			return
		}

		// remove bearer
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		// conver key
		decodedAccessPublicKey, err := base64.StdEncoding.DecodeString(m.container.Env.Auth.AccessTokenPublicKey)
		if err != nil {
			m.container.Logger.Error("api-middleware-DecodeString", err.Error())
			ctx.JSON(http.StatusUnauthorized, responses.ApiResponse{
				Status:  false,
				Message: "Could not decode key: " + err.Error(),
			})
			ctx.Abort()
			return
		}

		AccessPublicKey, err := jwt.ParseRSAPublicKeyFromPEM(decodedAccessPublicKey)
		if err != nil {
			m.container.Logger.Error("api-middleware-ParseRSAPublicKeyFromPEM", err.Error())
			ctx.JSON(http.StatusUnauthorized, responses.ApiResponse{
				Status:  false,
				Message: "Could not decode key: " + err.Error(),
			})
			ctx.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return AccessPublicKey, nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, responses.ApiResponse{
				Status:  false,
				Message: "Invalid access token: " + err.Error(),
			})
			ctx.Abort()
			return
		}

		// Check if token claims exist and have the expected format
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			m.container.Logger.Error("api-middleware-MapClaims", "Invalid token claims format! Token payload:", token.Claims)
			ctx.JSON(http.StatusUnauthorized, responses.ApiResponse{
				Status:  false,
				Message: "Invalid token claims format!",
			})
			ctx.Abort()
			return
		}

		referenceID, ok := claims["referenceId"].(uint)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, responses.ApiResponse{
				Status:  false,
				Message: "Invalid referenceId format!",
			})
			ctx.Abort()
			return
		}

		userType, ok := claims["userType"].(commonType.UserType)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, responses.ApiResponse{
				Status:  false,
				Message: "Invalid userType format!",
			})
			ctx.Abort()
			return
		}

		ctx.Set("ReferenceID", int(referenceID))
		ctx.Set("UserType", int(userType))
		ctx.Next()
	}
}

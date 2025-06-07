package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

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

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return m.container.Env.Auth.AccessPublicKey, nil
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

		// ✅ Check token expiration manually
		exp, ok := claims["exp"].(float64)
		if !ok || int64(exp) < time.Now().Unix() {
			ctx.JSON(http.StatusUnauthorized, responses.ApiResponse{
				Status:  false,
				Message: "Access token expired!",
			})
			ctx.Abort()
			return
		}

		ctx.Set("user_id", claims["user_id"])
		ctx.Set("user_type", claims["user_type"])
		ctx.Set("session_id", claims["session_id"])
		ctx.Next()
	}
}

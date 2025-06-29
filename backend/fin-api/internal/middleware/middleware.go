package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Uttamnath64/arvo-fin/app/common"
	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/requests"
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
	return func(c *gin.Context) {

		// ⏳ Create a context with timeout (e.g., 5 seconds)
		ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
		defer cancel()

		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, responses.ApiResponse{
				Status:  false,
				Message: "Missing access token.",
			})
			c.Abort()
			return
		}

		// remove bearer
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return m.container.Env.Auth.AccessPublicKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, responses.ApiResponse{
				Status:  false,
				Message: "Invalid or expired access token.",
			})
			c.Abort()
			return
		}

		// Check if token claims exist and have the expected format
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			m.container.Logger.Error("middleware-claims-format", "token", token.Raw)
			c.JSON(http.StatusUnauthorized, responses.ApiResponse{
				Status:  false,
				Message: "Invalid token claims format.",
			})
			c.Abort()
			return
		}

		// ✅ Check token expiration manually
		exp, ok := claims["exp"].(float64)
		if !ok || int64(exp) < time.Now().Unix() {
			c.JSON(http.StatusUnauthorized, responses.ApiResponse{
				Status:  false,
				Message: "Access token has expired.",
			})
			c.Abort()
			return
		}

		userIDFloat, ok := claims[string(common.CtxUserID)].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, responses.ApiResponse{
				Status:  false,
				Message: "Invalid user ID in token.",
			})
			c.Abort()
			return
		}

		sessionIDFloat, ok := claims[string(common.CtxSessionID)].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, responses.ApiResponse{
				Status:  false,
				Message: "Invalid session ID in token.",
			})
			c.Abort()
			return
		}

		userTypeFloat, ok := claims[string(common.CtxUserType)].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, responses.ApiResponse{
				Status:  false,
				Message: "Invalid user type in token.",
			})
			c.Abort()
			return
		}

		rctx := &requests.RequestContext{
			Ctx:       ctx,
			UserID:    uint(userIDFloat),
			UserType:  commonType.UserType(int8(userTypeFloat)),
			SessionID: uint(sessionIDFloat),
		}

		c.Set("rctx", rctx)
		c.Next()
	}
}

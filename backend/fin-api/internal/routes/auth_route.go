package routes

import (
	"github.com/Uttamnath64/arvo-fin/fin-api/internal/handlers"
	"github.com/Uttamnath64/arvo-fin/fin-api/internal/middleware"
)

func (routes *Routes) AuthRoutes() {
	handler := handlers.NewAuth(routes.container)
	middle := middleware.New(routes.container)
	userGroup := routes.server.Group("/auth").Use(middle.AuthMiddleware())
	{
		userGroup.POST("/login", handler.Login)
		userGroup.POST("/token", handler.Token)

		userGroup.POST("/send-otp", handler.SendOTP)
		userGroup.POST("/register", handler.Register)
		userGroup.POST("/reset-password", handler.ResetPassword)
	}
}

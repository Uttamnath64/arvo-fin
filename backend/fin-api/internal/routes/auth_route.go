package routes

import "github.com/Uttamnath64/arvo-fin/fin-api/internal/handlers"

func (routes *Routes) AuthRoutes() {
	handler := handlers.NewAuth(routes.container)
	userGroup := routes.server.Group("/auth")
	{
		userGroup.POST("/login", handler.Login)
		userGroup.POST("/register", handler.Register)
		userGroup.POST("/send-otp", handler.SentOTP)
		userGroup.POST("/reset-password", handler.ResetPassword)
		userGroup.POST("/token", handler.Token)
	}
}

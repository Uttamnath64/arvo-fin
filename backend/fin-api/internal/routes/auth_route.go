package routes

import "github.com/Uttamnath64/arvo-fin/fin-api/internal/handlers"

func (routes *Routes) AuthRoutes() {
	authHandler := handlers.NewAuthHandler(routes.container)
	userGroup := routes.server.Group("/auth")
	{
		userGroup.POST("/login", authHandler.LoginHandler)
		userGroup.POST("/register", authHandler.RegisterHandler)
		userGroup.POST("/send-otp", authHandler.SentOTPHandler)
		userGroup.POST("/reset-password", authHandler.ResetPasswordHandler)
	}
}

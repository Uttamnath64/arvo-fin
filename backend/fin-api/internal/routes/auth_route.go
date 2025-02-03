package routes

import "github.com/Uttamnath64/arvo-fin/fin-api/internal/handlers"

func (routes *Routes) AuthRoutes() {
	authHandler := handlers.NewAuthHandler(routes.container)
	userGroup := routes.server.Group("/auth")
	{
		userGroup.POST("/login", authHandler.LoginHandler)
		userGroup.POST("/register", authHandler.RegisterHandler)
		userGroup.POST("/get-otp", authHandler.GetOTPHandler)
		userGroup.POST("/forgot-password", authHandler.ForgotPasswordHandler)
		userGroup.POST("/reset-password", authHandler.ResetPasswordHandler)
	}
}

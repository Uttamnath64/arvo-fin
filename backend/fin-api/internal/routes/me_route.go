package routes

import "github.com/Uttamnath64/arvo-fin/fin-api/internal/handlers"

func (routes *Routes) MeRoutes() {
	handler := handlers.NewMe(routes.container)
	userGroup := routes.server.Group("/me")
	{
		userGroup.GET("/:user_id", handler.Get)
		userGroup.PUT("", handler.Update)
		userGroup.GET("/settings", handler.GetSettings)
		userGroup.PUT("/settings", handler.UpdateSettings)
	}
}

package routes

import "github.com/Uttamnath64/arvo-fin/fin-api/internal/handlers"

func (routes *Routes) MeRoutes() {
	handler := handlers.NewMe(routes.container)
	userGroup := routes.server.Group("/me")
	{
		userGroup.GET("/", handler.Get)
		userGroup.GET("/settings", handler.GetSettings)
		userGroup.PUT("", handler.Update)
		userGroup.PUT("/settings", handler.UpdateSettings)

		userGroup.GET("/:user_id", handler.Get)
		userGroup.GET("/settings/:user_id", handler.GetSettings)
	}
}

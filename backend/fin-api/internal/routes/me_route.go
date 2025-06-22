package routes

import (
	"github.com/Uttamnath64/arvo-fin/fin-api/internal/handlers"
	"github.com/Uttamnath64/arvo-fin/fin-api/internal/middleware"
)

func (routes *Routes) MeRoutes() {
	handler := handlers.NewMe(routes.container)
	middle := middleware.New(routes.container)
	userGroup := routes.server.Group("/me").Use(middle.Middleware())
	{
		userGroup.GET("/", handler.Get)
		userGroup.GET("/:userId", handler.Get)
		userGroup.GET("/settings/", handler.GetSettings)

		userGroup.PUT("/", handler.Update)
		userGroup.PUT("/settings/", handler.UpdateSettings)

	}
}

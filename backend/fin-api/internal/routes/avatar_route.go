package routes

import (
	"github.com/Uttamnath64/arvo-fin/fin-api/internal/handlers"
	"github.com/Uttamnath64/arvo-fin/fin-api/internal/middleware"
)

func (routes *Routes) avatarRoutes() {
	handler := handlers.NewAvatar(routes.container)
	middle := middleware.New(routes.container)
	userGroup := routes.server.Group("/avatar").Use(middle.Middleware())
	{
		userGroup.POST("/", handler.Create)
		userGroup.PUT("/:id", handler.Update)
		userGroup.GET("/:id", handler.Get)
		userGroup.GET("/type/:type", handler.GetAvatarsByType)
	}
}

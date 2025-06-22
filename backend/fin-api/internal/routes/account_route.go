package routes

import (
	"github.com/Uttamnath64/arvo-fin/fin-api/internal/handlers"
	"github.com/Uttamnath64/arvo-fin/fin-api/internal/middleware"
)

func (routes *Routes) AccountRoutes() {
	handler := handlers.NewAccount(routes.container)
	middle := middleware.New(routes.container)
	userGroup := routes.server.Group("/account").Use(middle.Middleware())
	{
		userGroup.GET("/types", handler.AccountTypes)

		userGroup.GET("/list/:portfolioId", handler.GetList)
		userGroup.GET("/:id", handler.Get)

		userGroup.POST("/", handler.Create)
		userGroup.PUT("/:id", handler.Update)
		userGroup.DELETE("/:id", handler.Delete)
	}
}

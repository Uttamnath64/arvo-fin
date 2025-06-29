package routes

import (
	"github.com/Uttamnath64/arvo-fin/fin-api/internal/handlers"
	"github.com/Uttamnath64/arvo-fin/fin-api/internal/middleware"
)

func (routes *Routes) PortfolioRoutes() {
	handler := handlers.NewPortfolio(routes.container)
	middle := middleware.New(routes.container)
	userGroup := routes.server.Group("/portfolio").Use(middle.Middleware())
	{
		userGroup.GET("/", handler.GetList)
		userGroup.GET("/user/:userId", handler.GetList)

		userGroup.GET("/:id", handler.Get)
		userGroup.POST("/", handler.Create)

		userGroup.PUT("/:id", handler.Update)
		userGroup.DELETE("/:id", handler.Delete)
	}
}

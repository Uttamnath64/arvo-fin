package routes

import (
	"github.com/Uttamnath64/arvo-fin/fin-api/internal/handlers"
	"github.com/Uttamnath64/arvo-fin/fin-api/internal/middleware"
)

func (routes *Routes) PortfolioRoutes() {
	portfolioHandler := handlers.NewPortfolio(routes.container)
	middle := middleware.New(routes.container)
	userGroup := routes.server.Group("/portfolio").Use(middle.Middleware())
	{
		userGroup.GET("", portfolioHandler.GetList)
		userGroup.GET("/:id", portfolioHandler.Get)
		userGroup.POST("", portfolioHandler.Add)
		userGroup.PUT("/:id", portfolioHandler.Update)
		userGroup.DELETE("/:id", portfolioHandler.Delete)
	}
}

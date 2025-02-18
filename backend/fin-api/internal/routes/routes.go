package routes

import (
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	container *storage.Container
	server    *gin.Engine
}

func New(container *storage.Container, server *gin.Engine) *Routes {
	return &Routes{
		container: container,
		server:    server,
	}
}

func (routes *Routes) Handlers() {
	routes.AuthRoutes()
	routes.UserRoutes()
	routes.PortfolioRoutes()
}

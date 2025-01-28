package routes

import (
	"context"

	"github.com/Uttamnath64/arvo-fin/app/config"
	"github.com/Uttamnath64/arvo-fin/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	ctx    context.Context
	server *gin.Engine
	config *config.Config
	logger *logger.Logger
}

func New(ctx context.Context, server *gin.Engine, con *config.Config, logger *logger.Logger) *Routes {
	return &Routes{
		ctx:    ctx,
		server: server,
		config: con,
		logger: logger,
	}
}

func (routes *Routes) Handlers() {
	routes.AuthRoutes()
	routes.UserRoutes()
}

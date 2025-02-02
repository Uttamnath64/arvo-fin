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
	env    *config.AppEnv
}

func New(ctx context.Context, server *gin.Engine, con *config.Config, logger *logger.Logger, env *config.AppEnv) *Routes {
	return &Routes{
		ctx:    ctx,
		server: server,
		config: con,
		logger: logger,
		env:    env,
	}
}

func (routes *Routes) Handlers() {
	routes.AuthRoutes()
	routes.UserRoutes()
}

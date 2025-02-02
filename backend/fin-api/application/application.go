package application

import (
	"context"
	"sync"
	"time"

	"github.com/Uttamnath64/arvo-fin/app/config"
	"github.com/Uttamnath64/arvo-fin/fin-api/internal/routes"
	"github.com/Uttamnath64/arvo-fin/pkg/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Application struct {
	wg    sync.WaitGroup
	state uint64 // set state
	// memcache *memcache.Memcache
	// cache *memory.MemoryCache // memory cache

	// main application context
	name   string
	ctx    context.Context
	cancel context.CancelFunc

	config *config.Config
	env    *config.AppEnv
	logger *logger.Logger

	// shutdownTimeout is the timeout for server. This timeout mean
	// that all component should be stopped after fetching signal
	// during this timeout.
	shutdownTimeout time.Duration
}

func New() *Application {
	return &Application{}
}

func (a *Application) Initialize() bool {
	var (
		con *config.Config
		err error
	)

	a.ctx = context.Background()

	// load env
	a.env, err = config.LoadEnv()
	if err != nil {
		logger.New("none").Error("api-application-env", err.Error())
		return false
	}
	a.logger = logger.New(a.env.Server.Environment)

	// load config DB
	err = config.LoadConfig(*a.env, con)
	if err != nil {
		a.logger.Error("api-application-config", err.Error())
		return false
	}
	a.config = con

	return true

}

func (a *Application) Run() {

	server := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{a.env.Server.ClientOrigin}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	// routers
	routes.New(a.ctx, server, a.config, a.logger, a.env).Handlers()

	if err := server.Run(":" + a.env.Server.Port); err != nil {
		a.logger.Error("api-application-server", err.Error())
		return
	}
}

func (a *Application) Name() string {
	return a.name
}

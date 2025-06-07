package application

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/Uttamnath64/arvo-fin/app/config"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/storage"
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
	cancel context.CancelFunc

	// di container
	container *storage.Container

	// shutdownTimeout is the timeout for server. This timeout mean
	// that all component should be stopped after fetching signal
	// during this timeout.
	shutdownTimeout time.Duration
}

func New() *Application {
	return &Application{}
}

func (a *Application) Initialize() bool {
	var con config.Config

	ctx := context.Background()
	requests.NewResponse()

	// load env
	env, err := config.LoadEnv(".env")
	if err != nil {
		logger.New("none").Error("api-application-env", err.Error())
		return false
	}

	// set logger
	log := logger.New(env.Server.Environment)

	// Load access and refresh keys
	err = config.LoadAccessAndRefreshKeys(&env)
	if err != nil {
		log.Error("api-application-accessAndRefreshKeys", err.Error())
		return false
	}

	// load config DB
	err = config.LoadConfig(env, &con)
	if err != nil {
		log.Error("api-application-config", err.Error())
		return false
	}

	// load redis
	redis, err := storage.NewRedisClient(ctx, env.Server.RedisAddr, "", 0)
	if err != nil {
		log.Error("api-application-redis", err.Error())
		return false
	}

	a.container = storage.NewContainer(ctx, &con, log, redis, &env)

	return true

}

func (a *Application) Run() {

	server := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{a.container.Env.Server.ClientOrigin}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	// routers
	routes.New(a.container, server).Handlers()
	if err := server.Run(fmt.Sprint(":", a.container.Env.Server.Port)); err != nil {
		a.container.Logger.Error("api-application-server", err.Error())
		return
	}
}

func (a *Application) Name() string {
	return a.name
}

package storage

import (
	"context"

	"github.com/Uttamnath64/arvo-fin/app/config"
	"github.com/Uttamnath64/arvo-fin/pkg/logger"
)

type Container struct {
	Ctx    context.Context
	Config *config.Config
	Logger *logger.Logger
	Redis  *RedisClient
	Env    *config.AppEnv
}

// NewContainer initializes the DI container
func NewContainer(ctx context.Context, cfg *config.Config, log *logger.Logger, redis *RedisClient, env *config.AppEnv) *Container {
	return &Container{
		Ctx:    ctx,
		Config: cfg,
		Logger: log,
		Redis:  redis,
		Env:    env,
	}
}

package services

import (
	"context"
	"os"

	"github.com/Uttamnath64/arvo-fin/app/config"
	"github.com/Uttamnath64/arvo-fin/app/requests"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"github.com/Uttamnath64/arvo-fin/pkg/logger"
)

func getTestContainer() (*storage.Container, bool) {
	var con config.Config

	ctx := context.Background()
	requests.NewResponse()

	// load env
	_ = os.Chdir("../../../")
	env, err := config.LoadEnv(".env.test")
	if err != nil {
		logger.New("none").Error("api-test-application-env", err.Error())
		return nil, false
	}

	// set logger
	log := logger.New(env.Server.Environment)

	// load redis
	redis, err := storage.NewRedisClient(ctx, env.Server.RedisAddr, "", 0)
	if err != nil {
		log.Error("api-test-application-redis", err.Error())
		return nil, false
	}

	return storage.NewContainer(ctx, &con, log, redis, &env), true
}

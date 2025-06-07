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
	var redis *storage.RedisClient
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
	log := logger.NewTest(env.Server.Environment)

	err = config.LoadAccessAndRefreshKeys(&env)
	if err != nil {
		log.Error("api-test-application-accessAndRefreshKeys", err.Error())
		return nil, false
	}

	return storage.NewContainer(ctx, &con, log, redis, &env), true
}

package main

import (
	"context"
	"os"

	"github.com/Uttamnath64/arvo-fin/app/config"
	script "github.com/Uttamnath64/arvo-fin/app/migrations/scripts"
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"github.com/Uttamnath64/arvo-fin/pkg/logger"
)

func getContainer() (*storage.Container, error) {
	var err error
	var con config.Config
	ctx := context.Background()

	// load env
	env, err := config.LoadEnv(".env")
	if err != nil {
		return nil, err
	}

	log := logger.New(env.Server.Environment)

	// load config DB
	err = config.LoadConfig(env, &con)
	if err != nil {
		return nil, err
	}

	// load redis
	redis, err := storage.NewRedisClient(ctx, env.Server.RedisAddr, "", 0)
	if err != nil {
		return nil, err
	}

	return storage.NewContainer(ctx, &con, log, redis, &env), nil
}

func main() {
	container, err := getContainer()
	if err != nil {
		logger.New("none").Error("api-application-env", err.Error())
	}

	// migration database
	err = container.Config.ReadWriteDB.AutoMigrate(
		&models.User{},
		&models.Portfolio{},
		&models.Account{},
		&models.Category{},
		&models.Budget{},
		&models.Avatar{},
		&models.Transaction{},
		&models.Session{},
		&models.RecurringTransaction{},
		&models.MigrationVersion{},
		&models.Currency{},
		&models.Admin{},
	)
	if err != nil {
		container.Logger.Error("app-migrate-config-error", "Failed to migrate the database", err)
		return
	}

	// migration log database
	err = container.Config.LogDB.AutoMigrate(
		&models.TransactionAudit{},
	)
	if err != nil {
		container.Logger.Error("app-migrate-config-error", "Failed to migrate the log database", err)
		return
	}

	if err := script.RunMigrations(container); err != nil {
		os.Exit(1)
	}

	container.Logger.Info("app-migrate-done", "👍 Migration completed!")
}

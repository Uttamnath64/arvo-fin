package main

import (
	"github.com/Uttamnath64/arvo-fin/app/config"
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/pkg/logger"
)

var (
	con config.Config
	env config.AppEnv
	log *logger.Logger
)

func init() {
	var err error

	// load env
	env, err = config.LoadEnv(".env")
	if err != nil {
		logger.New("none").Error("app-migrate-config-env", err.Error())
		return
	}

	log = logger.New(env.Server.Environment)

	// load config DB
	err = config.LoadConfig(env, &con)
	if err != nil {
		log.Error("app-migrate-config", err.Error())
		return
	}
}

func main() {
	var err error
	// migration database
	err = con.ReadWriteDB.AutoMigrate(
		&models.User{},
		&models.Portfolio{},
		&models.Account{},
		&models.Category{},
		&models.Budget{},
		&models.Avatar{},
		&models.Transaction{},
		&models.Session{},
		&models.RecurringTransaction{},
	)
	if err != nil {
		log.Error("app-migrate-config-error", "Failed to migrate the database")
		return
	}

	// migration log database
	err = con.LogDB.AutoMigrate(
		&models.TransactionAudit{},
	)
	if err != nil {
		log.Error("app-migrate-config-error", "Failed to migrate the log database")
		return
	}

	log.Info("app-migrate-done", "👍 Migration completed!")
}

package config

import (
	"database/sql"
	"encoding/base64"
	"fmt"

	"github.com/Uttamnath64/arvo-fin/pkg/logger"
	"github.com/golang-jwt/jwt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	ReadWriteDB *gorm.DB
	ReadOnlyDB  *gorm.DB
	LogDB       *gorm.DB
}

func EnsureDatabases(env *AppEnv, log *logger.Logger) error {

	db, err := sql.Open("mysql", env.Database.DSNMain)
	if err != nil {
		return fmt.Errorf("failed to connect to MySQL without DB: %w", err)
	}
	defer db.Close()

	queries := []string{
		fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`", env.Database.DB.MainDB),
		fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`", env.Database.DB.LogDB),
	}

	for _, query := range queries {
		log.Info("app-db-check", "running", query)
		if _, err := db.Exec(query); err != nil {
			return fmt.Errorf("error executing '%s': %w", query, err)
		}
	}

	log.Info("app-db-check", "message", "✅ Required databases ensured.")
	return nil
}

func LoadConfig(env AppEnv, con *Config) (err error) {
	con.ReadWriteDB, err = connect(env.Database.DSNReadWrite)
	if err != nil {
		return
	}
	con.ReadWriteDB = con.ReadWriteDB.Debug()

	con.ReadOnlyDB, err = connect(env.Database.DSNReadOnly)
	if err != nil {
		return
	}
	con.ReadOnlyDB = con.ReadOnlyDB.Debug()

	con.LogDB, err = connect(env.Database.DSNLogs)
	if err != nil {
		return
	}
	con.LogDB = con.LogDB.Debug()
	return
}

func LoadAccessAndRefreshKeys(env *AppEnv) error {
	// AccessPublicKey
	decodedAccessPublicKey, err := base64.StdEncoding.DecodeString(env.Auth.AccessTokenPublicKey)
	if err != nil {
		return err
	}
	env.Auth.AccessPublicKey, err = jwt.ParseRSAPublicKeyFromPEM(decodedAccessPublicKey)
	if err != nil {
		return err
	}

	// AccessPrivateKey
	decodedAccessPrivateKey, err := base64.StdEncoding.DecodeString(env.Auth.AccessTokenPrivateKey)
	if err != nil {
		return err
	}
	env.Auth.AccessPrivateKey, err = jwt.ParseRSAPrivateKeyFromPEM(decodedAccessPrivateKey)
	if err != nil {
		return err
	}

	// RefreshPublicKey
	decodedRefreshPublicKey, err := base64.StdEncoding.DecodeString(env.Auth.RefreshTokenPublicKey)
	if err != nil {
		return err
	}
	env.Auth.RefreshPublicKey, err = jwt.ParseRSAPublicKeyFromPEM(decodedRefreshPublicKey)
	if err != nil {
		return err
	}

	// RefreshPrivateKey
	decodedRefreshPrivateKey, err := base64.StdEncoding.DecodeString(env.Auth.RefreshTokenPrivateKey)
	if err != nil {
		return err
	}
	env.Auth.RefreshPrivateKey, err = jwt.ParseRSAPrivateKeyFromPEM(decodedRefreshPrivateKey)
	if err != nil {
		return err
	}
	return nil
}

// connect to DB
func connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

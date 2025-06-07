package config

import (
	"encoding/base64"

	"github.com/golang-jwt/jwt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	ReadWriteDB *gorm.DB
	ReadOnlyDB  *gorm.DB
	LogDB       *gorm.DB
}

func LoadConfig(env AppEnv, con *Config) (err error) {
	con.ReadWriteDB, err = connect(env.Database.DSNReadWrite)
	if err != nil {
		return
	}

	con.ReadOnlyDB, err = connect(env.Database.DSNReadOnly)
	if err != nil {
		return
	}

	con.LogDB, err = connect(env.Database.DSNLogs)
	if err != nil {
		return
	}
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

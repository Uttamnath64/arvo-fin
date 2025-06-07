package config

import (
	"crypto/rsa"
	"time"

	"github.com/spf13/viper"
)

// env config
type AppEnv struct {
	Database struct {
		DSNReadWrite string `mapstructure:"DSN_READ_WRITE"`
		DSNReadOnly  string `mapstructure:"DSN_READ_ONLY"`
		DSNLogs      string `mapstructure:"DSN_LOGS_DB"`
	}

	Server struct {
		Port         int    `mapstructure:"PORT"`
		ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
		Environment  string `mapstructure:"ENVIRONMENT"`
		RedisAddr    string `mapstructure:"REDIS_ADDR"`
		IsLive       bool   `mapstructure:"IS_LIVE"`
		Smtp         struct {
			Host     string `mapstructure:"SMTP_HOST"`
			Port     int    `mapstructure:"SMTP_PORT"`
			Email    string `mapstructure:"SMTP_EMAIL"`
			Password string `mapstructure:"SMTP_PASSWORD"`
		}
	}

	Auth struct {
		AccessTokenPublicKey   string `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
		AccessTokenPrivateKey  string `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
		RefreshTokenPublicKey  string `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
		RefreshTokenPrivateKey string `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
		AccessPublicKey        *rsa.PublicKey
		AccessPrivateKey       *rsa.PrivateKey
		RefreshPublicKey       *rsa.PublicKey
		RefreshPrivateKey      *rsa.PrivateKey
		AccessTokenExpired     time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRED"`
		RefreshTokenExpired    time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRED"`
	}
}

func LoadEnv(fileName string) (env AppEnv, err error) {

	viper.SetConfigName(fileName)
	viper.SetConfigType("env")
	viper.AddConfigPath("app/config/env")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	configs := []interface{}{&env.Database, &env.Server, &env.Auth, &env.Server.Smtp}
	for _, config := range configs {
		err = viper.Unmarshal(config)
		if err != nil {
			return
		}
	}

	return
}

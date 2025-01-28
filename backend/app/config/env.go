package config

import (
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
		Port         string `mapstructure:"PORT"`
		ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
		Environment  string `mapstructure:"ENVIRONMENT"`
	}

	Auth struct {
		AccessTokenPrivateKey  string        `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
		AccessTokenPublicKey   string        `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
		RefreshTokenPrivateKey string        `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
		RefreshTokenPublicKey  string        `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
		AccessTokenExpired     time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRED"`
		RefreshTokenExpired    time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRED"`
	}
}

func LoadEnv() (env *AppEnv, err error) {
	viper.AddConfigPath("env/")
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(env)
	return
}

package config

import (
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

// connect to DB
func connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

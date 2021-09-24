package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ConfigDb struct {
	DbHost     string
	DbUser     string
	DbPassword string
	DbName     string
	DbPort     string
	DbSslMode  string
	DbTimezone string
}

func (config *ConfigDb) InitialDb(debug bool) *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		config.DbHost,
		config.DbUser,
		config.DbPassword,
		config.DbName,
		config.DbPort,
		config.DbSslMode,
		config.DbTimezone,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if debug {
		db.Logger = logger.Default.LogMode(logger.Info)
	}
	if err != nil {
		panic(err.Error())
	}
	return db
}

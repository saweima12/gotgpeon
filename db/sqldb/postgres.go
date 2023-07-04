package sqldb

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitPostgresDb(dbUrl string, mode string) (*gorm.DB, error) {
	// define loglevel
	logLevel := logger.Info
	if mode == "pro" {
		logLevel = logger.Silent
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logLevel,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{Logger: newLogger})
	if err != nil {
		return nil, err
	}

	return db, nil
}

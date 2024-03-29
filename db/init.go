package db

import (
	"gotgpeon/config"
	"gotgpeon/db/cachedb"
	"gotgpeon/db/sqldb"
	"gotgpeon/logger"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	cache *redis.Client
	db    *gorm.DB
)

func InitDbConn(cfg *config.CommonConfig) error {
	// Set cache redis conenction.
	redisConn, err := cachedb.InitRedis(cfg.RedisUri)
	if err != nil {
		logger.Errorf("Redis connection err: %s, uri: %s", err.Error(), cfg.RedisUri)
	}
	cache = redisConn

	// Set database connection
	dbConn, err := sqldb.InitPostgresDb(cfg.DBUri)
	if err != nil {
		logger.Errorf("Database connection err: %s, uri: %s", err.Error(), cfg.DBUri)
		return err
	}
	db = dbConn

	return nil
}

func GetDB() *gorm.DB {
	return db
}

func GetCache() *redis.Client {
	return cache
}

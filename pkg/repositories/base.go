package repositories

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type BaseRepository struct {
	DbConn    *gorm.DB
	RedisConn *redis.Client
}

func (repo *BaseRepository) GetDB() *gorm.DB {
	return repo.DbConn
}

func (repo *BaseRepository) GetRedis() *redis.Client {
	return repo.RedisConn
}

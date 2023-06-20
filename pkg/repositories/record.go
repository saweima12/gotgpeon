package repositories

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type RecordRepository interface {
}

type recordRepository struct {
	BaseRepository
}

func NewRecordRepository(dbConn *gorm.DB, cacheConn *redis.Client) RecordRepository {
	return &recordRepository{
		BaseRepository: BaseRepository{DbConn: dbConn, RedisConn: cacheConn},
	}
}

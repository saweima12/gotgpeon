package repositories

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type BotConfigRepositroy struct {
	BaseRepository
}

func NewBotConfigRepo(dbConn *gorm.DB, redisConn *redis.Client) *BotConfigRepositroy {
	return &BotConfigRepositroy{
		BaseRepository: BaseRepository{DbConn: dbConn, RedisConn: redisConn},
	}
}

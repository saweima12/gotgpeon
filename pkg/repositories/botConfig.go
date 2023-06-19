package repositories

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type BotConfigRepository interface {
	GetChatConfig(chatId string)
}

func NewBotConfigRepo(dbConn *gorm.DB, redisConn *redis.Client) BotConfigRepository {
	return &botConfigRepositroy{
		BaseRepository: BaseRepository{DbConn: dbConn, RedisConn: redisConn},
	}
}

type botConfigRepositroy struct {
	BaseRepository
}

func (repo *botConfigRepositroy) GetChatConfig(chatId string) {
	panic("not implemented") // TODO: Implement
}

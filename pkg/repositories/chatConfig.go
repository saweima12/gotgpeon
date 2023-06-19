package repositories

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ChatConfigRepository interface {
}

type chatConfigRepository struct {
	BaseRepository
}

func NewChatConfigRepo(dbConn *gorm.DB, redisConn *redis.Client) ChatConfigRepository {
	return &chatConfigRepository{
		BaseRepository: BaseRepository{DbConn: dbConn, RedisConn: redisConn},
	}
}

func (repo *chatConfigRepository) GetChatConfig(chatId string) {
}

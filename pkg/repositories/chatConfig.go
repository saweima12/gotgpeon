package repositories

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ChatConfigRepository struct {
	BaseRepository
}

func NewChatConfigRepo(dbConn *gorm.DB, redisConn *redis.Client) *ChatConfigRepository {
	return &ChatConfigRepository{
		BaseRepository: BaseRepository{DbConn: dbConn, RedisConn: redisConn},
	}
}

func (repo *ChatConfigRepository) GetChatConfig(chatId string) {

}

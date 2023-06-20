package repositories

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type BotConfigRepository interface {
	GetWhiteList(chatId string)
	SetWhiteListCache(chatId string)
	SetWhiteListDB(chatId string)
	GetViolateRecord(chatId string, userId string)
	SetViolateCache(chatId string, userId string)
}

type botConfigRepository struct {
	BaseRepository
}

func NewBotConfigRepo(dbConn *gorm.DB, redisConn *redis.Client) BotConfigRepository {
	return &botConfigRepository{
		BaseRepository: BaseRepository{DbConn: dbConn, RedisConn: redisConn},
	}
}

func (repo *botConfigRepository) GetWhiteList(chatId string) {
	panic("not implemented") // TODO: Implement
}

func (repo *botConfigRepository) SetWhiteListCache(chatId string) {
	panic("not implemented") // TODO: Implement
}

func (repo *botConfigRepository) SetWhiteListDB(chatId string) {
	panic("not implemented") // TODO: Implement
}

func (botconfigrepositroy *botConfigRepository) GetViolateRecord(chatId string, userId string) {
	panic("not implemented") // TODO: Implement
}

func (botconfigrepositroy *botConfigRepository) SetViolateCache(chatId string, userId string) {
	panic("not implemented") // TODO: Implement
}

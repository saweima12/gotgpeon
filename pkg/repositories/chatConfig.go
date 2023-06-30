package repositories

import (
	"encoding/json"
	"gotgpeon/logger"
	"gotgpeon/models"
	"gotgpeon/models/entity"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ChatConfigRepository interface {
	GetChatConfig(chatId string) (*models.ChatConfig, error)
	SetConfigCache(chatId string, value *models.ChatConfig) error
	SetConfigDb(chatId string, value *models.ChatConfig)
}

type chatConfigRepository struct {
	BaseRepository
}

func NewChatConfigRepo(dbConn *gorm.DB, redisConn *redis.Client) ChatConfigRepository {
	return &chatConfigRepository{
		BaseRepository: BaseRepository{DbConn: dbConn, RedisConn: redisConn},
	}
}

// Load GroupChat Configuration.
func (repo *chatConfigRepository) GetChatConfig(chatId string) (*models.ChatConfig, error) {

	var result models.ChatConfig
	configKey := repo.getNameSpace(chatId)
	bytes, err := repo.GetRedis().Get(baseCtx, configKey).Bytes()

	// Redis has cache.
	if len(bytes) >= 1 && err == nil {
		err = json.Unmarshal(bytes, &result)
		if err != nil {
			logger.Errorf("GetChatConfig Unmarshal err: %s", err.Error())
			return nil, err
		}

		return &result, nil
	}
	// Redis don't have cache. Try load from database.
	configEntity := entity.PeonChatConfig{}
	err = repo.GetDB().Table(configEntity.TableName()).
		Select("*").
		Where("chat_id = ?", chatId).
		First(&configEntity).Error

	if err != nil {
		logger.Errorf("GetChatConfig DB err: %s", err.Error())
	}

	return &result, nil
}

func (repo *chatConfigRepository) SetConfigCache(chatId string, value *models.ChatConfig) error {
	// serialize to byte
	bytes, err := json.Marshal(value)
	if err != nil {
		logger.Errorf("SetConfigCache Marshal error: %s", err.Error())
	}

	key := repo.getNameSpace(chatId)
	err = repo.GetRedis().Set(baseCtx, key, bytes, 0).Err()
	if err != nil {
		logger.Errorf("SetConfigCache error: %s", err.Error())
		return err
	}

	return nil
}

func (repo *chatConfigRepository) SetConfigDb(chatId string, value *models.ChatConfig) {
	err := repo.GetDB().Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "chat_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"config_json", "attatch_json"}),
	}).Create(&value).Error

	if err != nil {
		logger.Errorf("SetConfigDb error: %s", err.Error())
		return
	}

}

func (repo *chatConfigRepository) getNameSpace(chatId string) string {
	return chatId + ":config"
}

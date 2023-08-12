package repositories

import (
	"encoding/json"
	"fmt"
	"gotgpeon/logger"
	"gotgpeon/models"
	"gotgpeon/models/entity"
	"strconv"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ChatRepository interface {
	GetAvaliableChatIds() []string
	GetChatConfig(chatId int64) (*models.ChatConfig, error)
	SetConfigCache(chatId int64, value *models.ChatConfig) error
	SetConfigDb(chatId int64, value *models.ChatConfig) error
	GetViolation(chatId int64, userId int64) (num int, err error)
	SetViolation(chatId int64, userId int64) (bool, error)
}

type chatRepository struct {
	BaseRepository
}

func NewChatRepo(dbConn *gorm.DB, redisConn *redis.Client) ChatRepository {
	return &chatRepository{
		BaseRepository: BaseRepository{DbConn: dbConn, RedisConn: redisConn},
	}
}

func (repo *chatRepository) GetAvaliableChatIds() []string {
	var result []string

	query := entity.PeonChatConfig{}
	err := repo.GetDB().Table(query.TableName()).
		Select("chat_id").Where("status = ?", 1).Find(&result).Error
	if err != nil {
		logger.Errorf("GetAvaliableChatIds err: %s", err.Error())
		return nil
	}

	return result
}

// Load GroupChat Configuration.
func (repo *chatRepository) GetChatConfig(chatId int64) (*models.ChatConfig, error) {

	var result *models.ChatConfig
	configKey := repo.getConfigNameSpace(chatId)

	rdb := repo.GetRedis()
	bytes, err := rdb.Get(baseCtx, configKey).Bytes()
	// Redis has cache.
	if len(bytes) >= 1 && err == nil {
		err = json.Unmarshal(bytes, &result)
		if err != nil {
			logger.Errorf("GetChatConfig Unmarshal err: %s", err.Error())
			return nil, err
		}
		return result, nil
	}

	// Redis don't have cache. Try load from database.
	configEntity := entity.PeonChatConfig{}
	err = repo.GetDB().Table(configEntity.TableName()).
		Select("*").
		Where("chat_id = ?", chatId).
		Take(&configEntity).Error

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(configEntity.ConfigJson, &result)
	if err != nil {
		return nil, err
	}

	repo.SetConfigCache(chatId, result)

	return nil, err
}

func (repo *chatRepository) SetConfigCache(chatId int64, value *models.ChatConfig) error {
	// serialize to byte
	bytes, err := json.Marshal(value)
	if err != nil {
		logger.Errorf("SetConfigCache Marshal error: %s", err.Error())
	}

	key := repo.getConfigNameSpace(chatId)
	err = repo.GetRedis().Set(baseCtx, key, bytes, 0).Err()
	if err != nil {
		logger.Errorf("SetConfigCache error: %s", err.Error())
		return err
	}
	return nil
}

func (repo *chatRepository) SetConfigDb(chatId int64, value *models.ChatConfig) error {

	// process parameter.
	bytes, err := json.Marshal(value)
	if err != nil {
		logger.Errorf("SetConfigDb Marshal error: %s", err.Error())
		return err
	}

	entityCfg := entity.PeonChatConfig{
		ChatId:     chatId,
		Status:     value.Status,
		ChatName:   value.ChatName,
		ConfigJson: bytes,
	}

	err = repo.GetDB().Model(&entityCfg).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "chat_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"status", "chat_name", "config_json"}),
	}).Create(&entityCfg).Error

	if err != nil {
		logger.Errorf("SetConfigDb error: %s", err.Error())
		return err
	}

	return nil
}

func (repo *chatRepository) GetViolation(chatId int64, userId int64) (num int, err error) {
	rdb := repo.GetRedis()
	key := repo.getViolationNamespace(chatId, userId)

	str, err := rdb.Get(baseCtx, key).Result()
	if err != nil {
		logger.Errorf("GetViolation error: %s", err.Error())
		return -1, err
	}

	num, err = strconv.Atoi(str)
	if err != nil {
		logger.Errorf("GetViolation value error: %s, value: %s", err.Error(), str)
		return -1, err
	}

	return num, nil
}

func (repo *chatRepository) SetViolation(chatId int64, userId int64) (bool, error) {
	// TODO: Didn't Implement.
	return false, nil
}

func (repo *chatRepository) getConfigNameSpace(chatId int64) string {
	return fmt.Sprintf("%d:config", chatId)
}

func (repo *chatRepository) getViolationNamespace(chatId int64, userId int64) string {
	return fmt.Sprintf("%d:violation:%d", chatId, userId)
}

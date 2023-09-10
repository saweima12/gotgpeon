package repositories

import (
	"fmt"
	"gotgpeon/data/entity"
	"gotgpeon/data/models"
	"gotgpeon/logger"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"gotgpeon/libs/json"
)

type ChatRepository interface {
	GetAvaliableChatIds() []int64
	GetChatCfg(chatId int64) (*models.ChatConfig, error)
	SetChatCfgCache(chatId int64, value *models.ChatConfig) error

	GetChatJobCfg(chatId int64) (*models.ChatJobConfig, error)
	SetChatJobCfgCache(chatId int64, value *models.ChatJobConfig) error
	UpdateChatCfgDB(chatId int64, newCfg *models.ChatConfig, newJobCfg *models.ChatJobConfig) error
}

type chatRepository struct {
	BaseRepository
}

func NewChatRepo(dbConn *gorm.DB, redisConn *redis.Client) ChatRepository {
	return &chatRepository{
		BaseRepository: BaseRepository{DbConn: dbConn, RedisConn: redisConn},
	}
}

func (repo *chatRepository) GetAvaliableChatIds() []int64 {
	var result []int64

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
func (repo *chatRepository) GetChatCfg(chatId int64) (*models.ChatConfig, error) {
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
	err = repo.GetDB().
		Where("chat_id = ?", chatId).
		Take(&configEntity).Error
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(configEntity.Config, &result)
	if err != nil {
		return nil, err
	}

	repo.SetChatCfgCache(chatId, result)
	return result, err
}

func (repo *chatRepository) SetChatCfgCache(chatId int64, value *models.ChatConfig) error {
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

func (repo *chatRepository) GetChatJobCfg(chatId int64) (*models.ChatJobConfig, error) {
	var result models.ChatJobConfig

	namespace := repo.getJobNamespace(chatId)
	// Try load config from redis
	byte, err := repo.GetRedis().Get(baseCtx, namespace).Bytes()
	if err == nil && len(byte) > 0 {
		err = json.Unmarshal(byte, &result)
		if err == nil {
			return &result, nil
		}
	}
	// Try load config from database.
	chatEntity := entity.PeonChatConfig{}
	err = repo.GetDB().Where("chat_id = ?", chatId).Take(&chatEntity).Error
	if err != nil {
		logger.Errorf("GetChatJobCfg getdb err: %s", err.Error())
		return nil, err
	}
	err = json.Unmarshal(chatEntity.JobConfig, &result)
	if err != nil {
		logger.Errorf("GetChatJobCfg db Unmarshal err: %s", err.Error())
		return nil, err
	}

	// get data success from database, write into redis.
	repo.SetChatJobCfgCache(chatId, &result)
	return nil, nil
}

func (repo *chatRepository) SetChatJobCfgCache(chatId int64, value *models.ChatJobConfig) error {

	bytes, err := json.Marshal(value)
	if err != nil {
		logger.Errorf("SetChatJobCfgCache marshal err: %s", err.Error())
		return err
	}

	namespace := repo.getJobNamespace(chatId)
	err = repo.GetRedis().Set(baseCtx, namespace, bytes, 0).Err()
	if err != nil {
		logger.Errorf("SetChatJobCfgCache err: %s", err.Error())
		return err
	}
	return nil
}

func (repo *chatRepository) UpdateChatCfgDB(chatId int64, newCfg *models.ChatConfig, newJobCfg *models.ChatJobConfig) error {
	// process parameter.
	bytes, err := json.Marshal(newCfg)
	if err != nil {
		logger.Errorf("SetConfigDb Marshal cfg error: %s", err.Error())
		return err
	}

	jobBytes, err := json.Marshal(newJobCfg)
	if err != nil {
		logger.Errorf("SetConfigDb Marshal jobCfg error: %s", err.Error())
		return err
	}

	entityCfg := entity.PeonChatConfig{
		ChatId:    chatId,
		Status:    newCfg.Status,
		ChatName:  newCfg.ChatName,
		Config:    bytes,
		JobConfig: jobBytes,
	}

	err = repo.GetDB().Model(&entityCfg).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "chat_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"status", "chat_name", "config", "job_config"}),
	}).Create(&entityCfg).Error

	if err != nil {
		logger.Errorf("SetConfigDb error: %s", err.Error())
		return err
	}

	return nil
}

func (repo *chatRepository) getConfigNameSpace(chatId int64) string {
	return fmt.Sprintf("%d:config", chatId)
}

func (repo *chatRepository) getJobNamespace(chatId int64) string {
	return fmt.Sprintf("%d:job", chatId)
}

package repositories

import (
	"encoding/json"
	"gotgpeon/logger"
	"gotgpeon/models"
	"gotgpeon/models/entity"

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

func (repo *recordRepository) GetUserRecord(chatId string, userId string) (result *models.MessageRecord, err error) {
	rdb := repo.GetRedis()
	// Get chat's point cache.
	namespace := getRecordNamespace(chatId)

	resp, err := rdb.HGet(baseCtx, namespace, userId).Bytes()
	if err == nil {
		// Cache has user's record.
		err := json.Unmarshal(resp, &result)
		if err != nil {
			logger.Errorf("GetUserRecord Unmarshal err: %s", err.Error())
			return nil, err
		}
		return result, nil
	}
	// Cache don't have user's record, try to find from database.
	entity := entity.PeonBehaviorRecord{}

	db := repo.GetDB()
	err = db.Select("point").Where("user_id = ?", userId).Take(&entity).Error
	if err != nil {
		logger.Errorf("GetUserRecord query err: %s", err.Error())
		return nil, err
	}

	result = &models.MessageRecord{
		UserId:      entity.UserId,
		Point:       entity.MsgCount,
		MemberLevel: entity.MemberLevel,
		CreatedTime: entity.CreatedTime,
	}

	// save to cache.

	return result, nil
}

func (repo *recordRepository) SetUserRecordCache(chatId string, record *models.MessageRecord) error {
	return nil
}

func (repo *recordRepository) SetUserRecordDB(chatId string, record *models.MessageRecord) error {
	return nil
}

func getRecordNamespace(chatId string) string {
	return chatId + "record_point"
}

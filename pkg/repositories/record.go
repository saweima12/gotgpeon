package repositories

import (
	"encoding/json"
	"errors"
	"gotgpeon/logger"
	"gotgpeon/models"
	"gotgpeon/models/entity"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RecordRepository interface {
	GetUserRecord(chatId string, query *models.MessageRecord) (result *models.MessageRecord, err error)
	SetUserRecordCache(chatId string, record *models.MessageRecord) error
	SetUserRecordDB(chatId string, record *models.MessageRecord) error
}

type recordRepository struct {
	BaseRepository
}

func NewRecordRepository(dbConn *gorm.DB, cacheConn *redis.Client) RecordRepository {
	return &recordRepository{
		BaseRepository: BaseRepository{DbConn: dbConn, RedisConn: cacheConn},
	}
}

func (repo *recordRepository) GetUserRecord(chatId string, query *models.MessageRecord) (result *models.MessageRecord, err error) {
	rdb := repo.GetRedis()
	// Get chat's point cache.
	namespace := getRecordNamespace(chatId)

	// Get Record struct
	resp, err := rdb.HGet(baseCtx, namespace, query.UserId).Bytes()
	if err == nil {
		// Cache has user's record.
		err := json.Unmarshal(resp, &result)
		if err != nil {
			logger.Errorf("GetUserRecord Unmarshal err: %s", err.Error())
			return nil, err
		}
		return result, nil
	}
	// Cache didn't have user record, try to find from database.
	entity := entity.PeonBehaviorRecord{}
	db := repo.GetDB()

	err = db.Where("user_id = ? AND chat_id = ?", query.UserId, chatId).
		Take(&entity).Error
	if err != nil && !errors.Is(gorm.ErrRecordNotFound, err) {
		logger.Errorf("GetUserRecord query db err: %v", err)
		return nil, err
	}

	result = &models.MessageRecord{
		UserId:      entity.UserId,
		FullName:    entity.FullName,
		MemberLevel: entity.MemberLevel,
		Point:       entity.MsgCount,
		CreatedTime: entity.CreatedTime,
	}

	repo.SetUserRecordCache(chatId, result)
	return result, nil

}

func (repo *recordRepository) SetUserRecordCache(chatId string, record *models.MessageRecord) error {
	rdb := repo.GetRedis()

	// Marshal to byte
	nameSpace := getRecordNamespace(chatId)
	byte, err := json.Marshal(record)
	if err != nil {
		logger.Errorf("SetUserRecordCache Marshal err: %s", err.Error())
		return err
	}

	// create insert map
	data := map[string]interface{}{}
	data[record.UserId] = byte

	// Try set cache
	err = rdb.HSet(baseCtx, nameSpace, data).Err()
	if err != nil {
		logger.Errorf("SetUserRecord err: %s", err.Error())
		return err
	}
	return nil
}

func (repo *recordRepository) SetUserRecordDB(chatId string, record *models.MessageRecord) error {

	schema := entity.PeonBehaviorRecord{
		ChatId:      chatId,
		UserId:      record.UserId,
		FullName:    record.FullName,
		MsgCount:    record.Point,
		MemberLevel: record.MemberLevel,
	}

	err := repo.GetDB().Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "chat_id"}, {Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"msg_count", "full_name"}),
	}).Create(&schema).Error
	if err != nil {
		logger.Errorf("SetUserRecordDB err:", err.Error())
		return err
	}

	return nil
}

func getRecordNamespace(chatId string) string {
	return chatId + ":record_point"
}

package repositories

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gotgpeon/logger"
	"gotgpeon/models"
	"gotgpeon/models/entity"
	"gotgpeon/utils/jsonutil"
	"strconv"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RecordRepository interface {
	GetAllUserRecordCache(chatId int64) (result map[int64]*models.MessageRecord, err error)
	GetUserRecord(chatId int64, query *models.MessageRecord) (result *models.MessageRecord, err error)
	GetAllCacheRecord(chatId int64) (result []models.MessageRecord, err error)
	SetUserRecordCache(chatId int64, record *models.MessageRecord) error
	SetUserRecordDB(chatId int64, record *models.MessageRecord) error
}

type recordRepository struct {
	BaseRepository
}

func NewRecordRepository(dbConn *gorm.DB, cacheConn *redis.Client) RecordRepository {
	return &recordRepository{
		BaseRepository: BaseRepository{DbConn: dbConn, RedisConn: cacheConn},
	}
}

func (repo *recordRepository) GetAllUserRecordCache(chatId int64) (result map[int64]*models.MessageRecord, err error) {
	rdb := repo.GetRedis()
	namespace := getRecordNamespace(chatId)

	result = make(map[int64]*models.MessageRecord)

	resp, err := rdb.HGetAll(baseCtx, namespace).Result()
	if err != nil {
		logger.Errorf("GetAllUserRecordCache err: %s", err.Error())
		return nil, err
	}
	for k, v := range resp {
		memberId, err := strconv.Atoi(k)
		memberId64 := int64(memberId)

		memberRecord := &models.MessageRecord{}
		err = jsonutil.UnmarshalFromString(v, memberRecord)
		if err != nil {
			logger.Errorf("GetAllUserRecordCache Unmarshal err: %s", err.Error())
			continue
		}
		result[memberId64] = memberRecord
	}

	return result, nil
}

func (repo *recordRepository) GetUserRecord(chatId int64, query *models.MessageRecord) (result *models.MessageRecord, err error) {
	rdb := repo.GetRedis()
	// Get chat's point cache.
	namespace := getRecordNamespace(chatId)

	// Get Record struct
	fieldKey := strconv.Itoa(int(query.MemberId))
	resp, err := rdb.HGet(baseCtx, namespace, fieldKey).Bytes()
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
	entity := entity.PeonChatMemberRecord{}
	db := repo.GetDB()

	err = db.Where("member_id = ? AND chat_id = ?", query.MemberId, chatId).
		Take(&entity).Error
	if err != nil {
		if !errors.Is(gorm.ErrRecordNotFound, err) {
			logger.Errorf("GetUserRecord query db err: %v", err)
		}
		return nil, err
	}

	return result, nil
}

func (repo *recordRepository) GetAllCacheRecord(chatId int64) (result []models.MessageRecord, err error) {
	// declare parameter
	result = make([]models.MessageRecord, 0)
	namespace := getRecordNamespace(chatId)

	resp, err := repo.GetRedis().HGetAll(context.Background(), namespace).Result()
	if err != nil {
		logger.Errorf("GetAllCacheRecord err: %s", err.Error())
		return nil, err
	}

	for _, value := range resp {
		tempRecord := models.MessageRecord{}
		err = jsonutil.UnmarshalFromString(value, &tempRecord)

		if err != nil {
			logger.Errorf("GetAllCacheRecord Unmarshal err: %s", err.Error())
			return nil, err
		}

		result = append(result, tempRecord)
	}
	return result, nil
}

func (repo *recordRepository) SetUserRecordCache(chatId int64, record *models.MessageRecord) error {
	rdb := repo.GetRedis()

	// Marshal to byte
	nameSpace := getRecordNamespace(chatId)
	byte, err := jsonutil.Marshal(record)
	if err != nil {
		logger.Errorf("SetUserRecordCache Marshal err: %s", err.Error())
		return err
	}

	// create insert map
	data := map[string]interface{}{}
	kMemberId := strconv.Itoa(int(record.MemberId))
	data[kMemberId] = byte

	// Try set cache
	err = rdb.HSet(baseCtx, nameSpace, data).Err()
	if err != nil {
		logger.Errorf("SetUserRecord err: %s", err.Error())
		return err
	}
	return nil
}

func (repo *recordRepository) SetUserRecordDB(chatId int64, record *models.MessageRecord) error {

	schema := entity.PeonChatMemberRecord{
		ChatId:      chatId,
		MemberId:    record.MemberId,
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

func getRecordNamespace(chatId int64) string {
	return fmt.Sprintf("%d:record_point", chatId)
}

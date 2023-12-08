package repositories

import (
	"encoding/json"
	"gotgpeon/data/entity"
	"gotgpeon/data/models"
	"gotgpeon/logger"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type DeletedMsgRepository interface {
	GetList(chatId int64) ([]*models.DeletedMessage, error)
	Insert(chatId int64, contentType string, jsonBytes []byte) error
	CleanOutdated() error
}

type deletedMsgRepository struct {
	BaseRepository
}

func NewDeletedMsgRepository(dbConn *gorm.DB, cacheConn *redis.Client) DeletedMsgRepository {
	return &deletedMsgRepository{
		BaseRepository: BaseRepository{DbConn: dbConn, RedisConn: cacheConn},
	}
}

func (repo *deletedMsgRepository) GetList(chatId int64) ([]*models.DeletedMessage, error) {

	var err error
	var entities []*entity.PeonDeletedMessage

	err = repo.GetDB().
		Where("chat_id = ?", chatId).
		Find(&entities).Error
	if err != nil {
		logger.Errorf("[deletedMsgRepository] GetList - err: %s", err.Error())
		return nil, err
	}

	result := []*models.DeletedMessage{}
	for eid := range entities {
		item := models.DeletedMessage{
			ContentType: entities[eid].ContentType,
			Content:     json.RawMessage(entities[eid].MessageJson),
			RecordTime:  entities[eid].RecordTime.Unix(),
		}
		result = append(result, &item)
	}

	return result, nil
}

func (repo *deletedMsgRepository) Insert(chatId int64, contentType string, jsonBytes []byte) error {
	newEntity := entity.PeonDeletedMessage{
		ChatId:      chatId,
		ContentType: contentType,
		MessageJson: jsonBytes,
	}

	err := repo.GetDB().Create(&newEntity).Error
	if err != nil {
		logger.Errorf("InsertDeletedRecord err: %s", err.Error())
		return err
	}
	return nil
}

func (repo *deletedMsgRepository) CleanOutdated() error {
	// process parameter
	now := time.Now().UTC()
	offset := now.AddDate(0, 0, -14)

	// query
	item := entity.PeonDeletedMessage{}
	err := repo.GetDB().Where("record_date < ?", offset).Delete(&item).Error
	if err != nil {
		logger.Errorf("DeleteOutdatedRecordList err: %s", err.Error())
		return err
	}
	return nil
}

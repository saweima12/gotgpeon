package repositories

import (
	"gotgpeon/logger"
	"gotgpeon/data/entity"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type DeletedMsgRepository interface {
	GetDeletedRecordListByChat(chatId int64) ([]*entity.PeonDeletedMessage, error)
	InsertDeletedRecord(chatId int64, contentType string, jsonBytes []byte) error
	DeleteOutdatedRecordList() error
}

type deletedMsgRepository struct {
	BaseRepository
}

func NewDeletedMsgRepository(dbConn *gorm.DB, cacheConn *redis.Client) DeletedMsgRepository {
	return &deletedMsgRepository{
		BaseRepository: BaseRepository{DbConn: dbConn, RedisConn: cacheConn},
	}
}

func (repo *deletedMsgRepository) GetDeletedRecordListByChat(chatId int64) (result []*entity.PeonDeletedMessage, err error) {
	err = repo.GetDB().Where("chat_id = ?", chatId).Find(&result).Error
	if err != nil {
		logger.Errorf("GetDeletedRecordListByChat err: %s", err.Error())
		return nil, err
	}

	return result, nil
}

func (repo *deletedMsgRepository) InsertDeletedRecord(chatId int64, contentType string, jsonBytes []byte) error {
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

func (repo *deletedMsgRepository) DeleteOutdatedRecordList() error {
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

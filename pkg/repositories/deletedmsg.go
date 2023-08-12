package repositories

import (
	"gotgpeon/models/entity"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type DeletedMsgRepository interface {
	InsertDeletedRecord(chatId int64, contentType string, jsonBytes []byte) error
}

type deletedMsgRepository struct {
	BaseRepository
}

func NewDeletedMsgRepository(dbConn *gorm.DB, cacheConn *redis.Client) DeletedMsgRepository {
	return &deletedMsgRepository{
		BaseRepository: BaseRepository{DbConn: dbConn, RedisConn: cacheConn},
	}
}

func (repo *deletedMsgRepository) InsertDeletedRecord(chatId int64, contentType string, jsonBytes []byte) error {
	newEntity := entity.PeonDeletedMessage{
		ChatId:      chatId,
		ContentType: contentType,
		MessageJson: jsonBytes,
	}

	err := repo.GetDB().Create(&newEntity).Error
	if err != nil {
		return err
	}

	return nil
}

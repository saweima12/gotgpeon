package services

import (
	"gotgpeon/libs/json"
	"gotgpeon/logger"
	"gotgpeon/models/entity"
	"gotgpeon/pkg/repositories"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type DeletedService interface {
	GetDeletedRecordListByChat(chatId int64) []*entity.PeonDeletedMessage
	InsertDeletedRecord(chatId int64, contentType string, message *tgbotapi.Message) error
	DeleteOutdatedRecordList() error
}

type deletedService struct {
	DeletedMsgRepo repositories.DeletedMsgRepository
}

func NewDeletedService(deletedMsgRepo repositories.DeletedMsgRepository) DeletedService {
	return &deletedService{
		DeletedMsgRepo: deletedMsgRepo,
	}
}

func (s *deletedService) GetDeletedRecordListByChat(chatId int64) []*entity.PeonDeletedMessage {
	records, err := s.DeletedMsgRepo.GetDeletedRecordListByChat(chatId)
	if err != nil {
		return []*entity.PeonDeletedMessage{}
	}
	return records
}

func (s *deletedService) DeleteOutdatedRecordList() error {
	return s.DeletedMsgRepo.DeleteOutdatedRecordList()
}

func (s *deletedService) InsertDeletedRecord(chatId int64, contentType string, message *tgbotapi.Message) error {
	jsonBytes, err := json.Marshal(message)
	if err != nil {
		logger.Errorf("InsertDeletedRecord marshal err: %s || msg: %v", err.Error(), message)
	}
	// Add to database.
	err = s.DeletedMsgRepo.InsertDeletedRecord(chatId, contentType, jsonBytes)
	if err != nil {
		logger.Errorf("InsertDeletedRecord err: %s", err.Error())
	}

	return nil
}

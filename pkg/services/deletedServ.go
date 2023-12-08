package services

import (
	"gotgpeon/data/models"
	"gotgpeon/libs/json"
	"gotgpeon/logger"
	"gotgpeon/pkg/repositories"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type DeletedService interface {
	GetListByChat(chatId int64) []*models.DeletedMessage
	Insert(chatId int64, contentType string, message *tgbotapi.Message) error
	CleanOutdated() error
}

type deletedService struct {
	DeletedMsgRepo repositories.DeletedMsgRepository
}

func NewDeletedService(deletedMsgRepo repositories.DeletedMsgRepository) DeletedService {
	return &deletedService{
		DeletedMsgRepo: deletedMsgRepo,
	}
}

func (s *deletedService) GetListByChat(chatId int64) []*models.DeletedMessage {
	records, err := s.DeletedMsgRepo.GetList(chatId)
	if err != nil {
		return []*models.DeletedMessage{}
	}
	return records
}

func (s *deletedService) CleanOutdated() error {
	return s.DeletedMsgRepo.CleanOutdated()
}

func (s *deletedService) Insert(chatId int64, contentType string, message *tgbotapi.Message) error {
	jsonBytes, err := json.Marshal(message)
	if err != nil {
		logger.Errorf("InsertDeletedRecord marshal err: %s || msg: %v", err.Error(), message)
	}
	// Add to database.
	err = s.DeletedMsgRepo.Insert(chatId, contentType, jsonBytes)
	if err != nil {
		logger.Errorf("InsertDeletedRecord err: %s", err.Error())
	}

	return nil
}

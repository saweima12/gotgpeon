package services

import (
	"gotgpeon/models"
	"gotgpeon/pkg/repositories"
)

type RecordService interface {
	GetUserRecord(chatId string, query *models.MessageRecord) *models.MessageRecord
	SetUserRecord(chatId string, record *models.MessageRecord) error
}

type recordService struct {
	RecordRepo repositories.RecordRepository
}

func NewRecordService(recordRepo repositories.RecordRepository) RecordService {
	return &recordService{
		RecordRepo: recordRepo,
	}
}

func (s recordService) GetUserRecord(chatId string, query *models.MessageRecord) *models.MessageRecord {
	record, err := s.RecordRepo.GetUserRecord(chatId, query)
	if err != nil {
		return models.NewMessageRecord(query.UserId, query.FullName)
	}
	return record
}

func (s recordService) SetUserRecord(chatId string, data *models.MessageRecord) error {
	err := s.RecordRepo.SetUserRecordCache(chatId, data)
	if err != nil {
		return err
	}

	err = s.RecordRepo.SetUserRecordDB(chatId, data)
	if err != nil {
		return err
	}
	return nil
}

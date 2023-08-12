package services

import (
	"gotgpeon/models"
	"gotgpeon/pkg/repositories"
)

type RecordService interface {
	GetUserRecord(chatId int64, query *models.MessageRecord) *models.MessageRecord
	SetUserRecordCache(chatId int64, record *models.MessageRecord) error
	SetUserRecordDB(chatId int64, record *models.MessageRecord) error
}

type recordService struct {
	RecordRepo repositories.RecordRepository
}

func NewRecordService(recordRepo repositories.RecordRepository) RecordService {
	return &recordService{
		RecordRepo: recordRepo,
	}
}

func (s recordService) GetUserRecord(chatId int64, query *models.MessageRecord) *models.MessageRecord {
	record, err := s.RecordRepo.GetUserRecord(chatId, query)

	if err != nil {
		return models.NewMessageRecord(query.MemberId, query.FullName)
	}

	return record
}

func (s recordService) SetUserRecordCache(chatId int64, data *models.MessageRecord) error {
	err := s.RecordRepo.SetUserRecordCache(chatId, data)
	if err != nil {
		return err
	}

	return nil
}

func (s recordService) SetUserRecordDB(chatId int64, data *models.MessageRecord) error {
	err := s.RecordRepo.SetUserRecordDB(chatId, data)
	if err != nil {
		return err
	}

	return nil
}

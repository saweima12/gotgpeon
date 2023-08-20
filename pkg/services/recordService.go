package services

import (
	"gotgpeon/logger"
	"gotgpeon/models"
	"gotgpeon/pkg/repositories"
	"strconv"
)

type RecordService interface {
	GetAllUserRecord(chatId int64) map[int64]*models.MessageRecord
	GetUserRecord(chatId int64, query *models.MessageRecord) *models.MessageRecord
	SetUserRecordCache(chatId int64, record *models.MessageRecord) error
	SetUserRecordDB(chatId int64, record *models.MessageRecord) error
	DelCacheByMemberIds(chatId int64, memberIdList []int64) error
}

type recordService struct {
	RecordRepo repositories.RecordRepository
}

func NewRecordService(recordRepo repositories.RecordRepository) RecordService {
	return &recordService{
		RecordRepo: recordRepo,
	}
}

func (s *recordService) GetAllUserRecord(chatId int64) map[int64]*models.MessageRecord {
	records, err := s.RecordRepo.GetAllUserRecordCache(chatId)
	if err != nil {
		logger.Errorf("GetAllUserRecord err: %s", err.Error())
		return nil
	}
	return records
}

func (s *recordService) GetUserRecord(chatId int64, query *models.MessageRecord) *models.MessageRecord {
	record, err := s.RecordRepo.GetUserRecord(chatId, query)

	if err != nil {
		record = models.NewMessageRecord(query.MemberId, query.FullName)
	}
	// overwirte fullname to latest.
	record.FullName = query.FullName

	return record
}

func (s *recordService) SetUserRecordCache(chatId int64, data *models.MessageRecord) error {
	err := s.RecordRepo.SetUserRecordCache(chatId, data)
	if err != nil {
		return err
	}

	return nil
}

func (s *recordService) SetUserRecordDB(chatId int64, data *models.MessageRecord) error {
	err := s.RecordRepo.SetUserRecordDB(chatId, data)
	if err != nil {
		return err
	}

	return nil
}

func (s *recordService) DelCacheByMemberIds(chatId int64, memberIdList []int64) error {
	strIds := make([]string, len(memberIdList))
	for _, id := range memberIdList {
		idStr := strconv.Itoa(int(id))
		strIds = append(strIds, idStr)
	}

	err := s.RecordRepo.DelCacheByMemberIds(chatId, strIds)
	if err != nil {
		return err
	}
	return nil
}

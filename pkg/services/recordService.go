package services

import "gotgpeon/models"

type RecordService interface{}

type recordService struct{}

func NewRecordService() RecordService {
	return nil
}

func (s recordService) GetRecord(chatId string, userId string) *models.MessageRecord {

}

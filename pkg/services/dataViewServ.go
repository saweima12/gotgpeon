package services

import (
	"gotgpeon/data/models"
	"gotgpeon/logger"
	"gotgpeon/pkg/repositories"
)

type DataviewService interface {
	GetChatList() map[int64]string
	GetMemberListByChat(chatId int64) []*models.ChatMemberResult
	GetDeletedMsgListByChat(chatId int64)
}

type dataviewService struct {
	chatRepo   repositories.ChatRepository
	recordRepo repositories.RecordRepository
}

func NewDataviewService(
	chatRepo repositories.ChatRepository,
	recordRepo repositories.RecordRepository,
) DataviewService {
	return &dataviewService{
		chatRepo:   chatRepo,
		recordRepo: recordRepo,
	}
}

func (da *dataviewService) GetChatList() map[int64]string {
	result, err := da.chatRepo.GetAvaliableChatList()
	if err != nil {
		logger.Error("GetChatList err:", err.Error())
		return map[int64]string{}
	}
	return result
}

func (da *dataviewService) GetMemberListByChat(chatId int64) []*models.ChatMemberResult {
	result, err := da.recordRepo.GetAllUserRecord(chatId)
	if err != nil {
		logger.Error("GetMemberListByChat err:", err.Error())
		return []*models.ChatMemberResult{}
	}
	return result
}

func (da *dataviewService) GetDeletedMsgListByChat(chatId int64) {
	panic("not implemented") // TODO: Implement
}

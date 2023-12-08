package services

import (
	"gotgpeon/data/models"
	"gotgpeon/logger"
	"gotgpeon/pkg/repositories"
)

type DataviewService interface {
	GetChatList() map[int64]string
	GetMemberListByChat(chatId int64) []*models.ChatMemberInfo
	GetDeletedMsgListByChat(chatId int64) []*models.DeletedMessage
}

type dataviewService struct {
	chatRepo    repositories.ChatRepository
	recordRepo  repositories.MemberRecordRepository
	deletedRepo repositories.DeletedMsgRepository
}

func NewDataviewService(
	chatRepo repositories.ChatRepository,
	recordRepo repositories.MemberRecordRepository,
	deletedRepo repositories.DeletedMsgRepository,
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

func (da *dataviewService) GetMemberListByChat(chatId int64) []*models.ChatMemberInfo {
	result, err := da.recordRepo.GetAllUserRecord(chatId)
	if err != nil {
		logger.Error("GetMemberListByChat err:", err.Error())
		return []*models.ChatMemberInfo{}
	}
	return result
}

func (da *dataviewService) GetDeletedMsgListByChat(chatId int64) []*models.DeletedMessage {
	list, err := da.deletedRepo.GetList(chatId)
	if err != nil {
		logger.Errorf("GetDeletedMsgListByChat err: %v", err)
		return []*models.DeletedMessage{}
	}

	return list
}

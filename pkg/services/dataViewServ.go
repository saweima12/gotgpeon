package services

import "gotgpeon/pkg/repositories"

type DataviewService interface {
	GetChatList()
	GetMemberListByChat(chatId int64)
	GetDeletedMsgListByChat(chatId int64)
}

type dataviewService struct {
	chatRepo repositories.ChatRepository
}

func NewDataviewService() DataviewService {
	return nil
}

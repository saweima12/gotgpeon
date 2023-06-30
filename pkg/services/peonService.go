package services

import (
	"gotgpeon/models"
	"gotgpeon/pkg/repositories"
)

type PeonService interface {
	GetChatConfig(chatId string, chatName string) *models.ChatConfig
	SetChatConfig(cfg *models.ChatConfig)
	GetBotAllowlist() map[string]byte
	IsAllowListUser(userId string) bool
}

type peonService struct {
	chatRepo repositories.ChatRepository
	botRepo  repositories.BotConfigRepository
}

func NewPeonService(chatRepo repositories.ChatRepository, botRepo repositories.BotConfigRepository) PeonService {
	return &peonService{
		chatRepo: chatRepo,
		botRepo:  botRepo,
	}
}

func (s peonService) GetChatConfig(chatId string, chatName string) *models.ChatConfig {

	chatCfg, err := s.chatRepo.GetChatConfig(chatId)
	if err != nil {
		chatCfg = models.NewDefaultChatConfig(chatId, chatName, []string{})
	}
	return chatCfg
}

func (s peonService) SetChatConfig(newCfg *models.ChatConfig) {
	chatId := newCfg.ChatId
	// Save to cache.
	s.chatRepo.SetConfigCache(chatId, newCfg)
	// Save to database.
	s.chatRepo.SetConfigDb(chatId, newCfg)
}

func (s *peonService) GetBotAllowlist() map[string]byte {
	allowList := s.botRepo.GetWhiteList()
	return allowList
}

func (s *peonService) IsAllowListUser(userId string) bool {
	allowList := s.GetBotAllowlist()
	if val, ok := allowList[userId]; ok {
		if val == 1 {
			return true
		}
	}
	return false
}

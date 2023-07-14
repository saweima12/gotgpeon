package services

import (
	"gotgpeon/logger"
	"gotgpeon/models"
	"gotgpeon/pkg/repositories"
)

type PeonService interface {
	GetChatConfig(chatId string, chatName string) *models.ChatConfig
	SetChatConfig(cfg *models.ChatConfig) error
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
		return models.NewDefaultChatConfig(chatId, chatName, []string{})
	}

	return chatCfg
}

func (s peonService) SetChatConfig(newCfg *models.ChatConfig) error {
	chatId := newCfg.ChatId
	// Save to database.
	err := s.chatRepo.SetConfigDb(chatId, newCfg)

	if err != nil {
		logger.Errorf("SetChatConfigDb err: %s", err.Error())
		return err
	}

	// Save to cache.
	s.chatRepo.SetConfigCache(chatId, newCfg)
	if err != nil {
		logger.Errorf("SetChatConfigCache err: %s", err.Error())
	}

	return nil
}

func (s *peonService) GetBotAllowlist() map[string]byte {
	return s.botRepo.GetWhiteList()
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

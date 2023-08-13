package services

import (
	"gotgpeon/logger"
	"gotgpeon/models"
	"gotgpeon/pkg/repositories"
	"gotgpeon/utils/jsonutil"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type PeonService interface {
	GetChatConfig(chatId int64, chatName string) *models.ChatConfig
	SetChatConfig(cfg *models.ChatConfig) error
	GetBotAllowlist() map[int64]byte
	IsAllowListUser(userId int64) bool
	InsertDeletedRecord(chatId int64, contentType string, message *tgbotapi.Message) error
}

type peonService struct {
	chatRepo       repositories.ChatRepository
	botRepo        repositories.BotConfigRepository
	deletedMsgRepo repositories.DeletedMsgRepository
}

func NewPeonService(
	chatRepo repositories.ChatRepository,
	botRepo repositories.BotConfigRepository,
	deletedMsgRepo repositories.DeletedMsgRepository,
) PeonService {
	return &peonService{
		chatRepo:       chatRepo,
		botRepo:        botRepo,
		deletedMsgRepo: deletedMsgRepo,
	}
}

func (s peonService) GetChatConfig(chatId int64, chatName string) *models.ChatConfig {
	chatCfg, err := s.chatRepo.GetChatConfig(chatId)

	if err != nil {
		return models.NewDefaultChatConfig(chatId, chatName, []int64{})
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

func (s *peonService) InsertDeletedRecord(chatId int64, contentType string, message *tgbotapi.Message) error {
	jsonBytes, err := jsonutil.Marshal(message)
	if err != nil {
		logger.Errorf("InsertDeletedRecord marshal err: %s || msg: %v", err.Error(), message)
	}
	// Add to database.
	err = s.deletedMsgRepo.InsertDeletedRecord(chatId, contentType, jsonBytes)
	if err != nil {
		logger.Errorf("InsertDeletedRecord err: %s", err.Error())
	}

	return nil
}

func (s *peonService) GetBotAllowlist() map[int64]byte {
	return s.botRepo.GetAllowlist()
}

func (s *peonService) IsAllowListUser(userId int64) bool {
	allowList := s.GetBotAllowlist()
	if val, ok := allowList[userId]; ok {
		if val == 1 {
			return true
		}
	}
	return false
}

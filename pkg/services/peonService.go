package services

import (
	"gotgpeon/logger"
	"gotgpeon/models"
	"gotgpeon/pkg/repositories"
	"gotgpeon/utils/jsonutil"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type PeonService interface {
	GetChatConfig(chatId int64) *models.ChatConfig
	SetChatConfig(cfg *models.ChatConfig) error
	GetChatJobConfig(chatId int64) *models.ChatJobConfig
	SetChatJobConfig(chatId int64, newJobCfg *models.ChatJobConfig) error

	UpdateChatConfigDB(chatId int64) error

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

func (s *peonService) GetChatConfig(chatId int64) *models.ChatConfig {
	chatCfg, err := s.chatRepo.GetChatCfg(chatId)
	if err != nil {
		return models.NewDefaultChatConfig(chatId, []int64{})
	}
	return chatCfg
}

func (s *peonService) SetChatConfig(newCfg *models.ChatConfig) error {
	// Save to cache.
	err := s.chatRepo.SetChatCfgCache(newCfg.ChatId, newCfg)
	if err != nil {
		logger.Errorf("SetChatConfigCache err: %s", err.Error())
	}
	return nil
}

func (s *peonService) GetChatJobConfig(chatId int64) *models.ChatJobConfig {
	jobCfg, err := s.chatRepo.GetChatJobCfg(chatId)
	if err != nil {
		return models.NewDefaultChatJobConfig()
	}
	return jobCfg
}

func (s *peonService) SetChatJobConfig(chatId int64, newJobCfg *models.ChatJobConfig) error {
	err := s.chatRepo.SetChatJobCfgCache(chatId, newJobCfg)
	if err != nil {
		logger.Errorf("SetChatJobConfig err: %s", err.Error())
	}
	return nil
}

func (s *peonService) UpdateChatConfigDB(chatId int64) error {
	newCfg := s.GetChatConfig(chatId)
	newJobCfg := s.GetChatJobConfig(chatId)
	s.chatRepo.UpdateChatCfgDB(chatId, newCfg, newJobCfg)

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

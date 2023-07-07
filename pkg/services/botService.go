package services

import (
	"gotgpeon/logger"
	"gotgpeon/models"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotService interface {
}

type botService struct {
	BotAPI *tgbotapi.BotAPI
}

func NewBotService() BotService {
	return &botService{}
}

func (s *botService) SendMessage(message tgbotapi.Chattable, duration time.Duration) {
	s.BotAPI.Send(message)

	if duration > 0 {
		// TODO: Delay delete message.
	}
}

func (s *botService) SetPermission(chatId int64, userId int64, level int, until_date int64) {
	permission := getPermissionByLevel(level)
	operate := tgbotapi.RestrictChatMemberConfig{
		ChatMemberConfig: tgbotapi.ChatMemberConfig{ChatID: chatId, UserID: userId},
		UntilDate:        until_date,
		Permissions:      permission,
	}

	_, err := s.BotAPI.Send(operate)
	if err != nil {
		logger.Errorf("SetPermission err: %s", err.Error())
	}

	// TODO: Send setPermission tips.
}

func getPermissionByLevel(level int) *tgbotapi.ChatPermissions {
	result := &tgbotapi.ChatPermissions{}

	if level >= models.NONE {
		result.CanSendMessages = true
	}

	if level >= models.JUNIOR {
		result.CanSendOtherMessages = true
		result.CanSendMediaMessages = true
		result.CanAddWebPagePreviews = true
	}

	return result
}

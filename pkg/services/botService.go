package services

import (
	"fmt"
	"gotgpeon/logger"
	"gotgpeon/models"
	"gotgpeon/pkg/tgbot/boterr"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotService interface {
	SendMessage(message tgbotapi.Chattable, duration time.Duration)
	DeleteMessage(chatId int64, messageId int)
}

type botService struct {
	BotAPI *tgbotapi.BotAPI
}

func NewBotService(botAPI *tgbotapi.BotAPI) BotService {
	return &botService{
		BotAPI: botAPI,
	}
}

func (s *botService) SendMessage(message tgbotapi.Chattable, duration time.Duration) {
	resultMsg, err := s.BotAPI.Send(message)

	if err != nil {
		logger.Errorf("Send message err: %v", err)
		return
	}

	if duration > 0 {
		// TODO: Delay delete message.
		fmt.Println(resultMsg)
	}
}

func (s *botService) DeleteMessage(chatId int64, messageId int) {
	deleteReq := tgbotapi.NewDeleteMessage(chatId, messageId)
	_, err := s.BotAPI.Request(deleteReq)

	if err != nil {
		if boterr.IsNotFound(err) {
			return
		}
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

package services

import (
	"gotgpeon/libs/timewheel"
	"gotgpeon/logger"
	"gotgpeon/models"
	"gotgpeon/pkg/tgbot/boterr"
	"time"

	"github.com/avast/retry-go"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type delayDeleteTask struct {
	MessageId  int
	ChatId     int64
	BotService BotService
}

func (t *delayDeleteTask) Run() {
	t.BotService.DeleteMessageById(t.ChatId, t.MessageId)
}

type BotService interface {
	SendMessage(message tgbotapi.Chattable, duration time.Duration)
	DeleteMessageById(chatId int64, messageId int)
	SetPermission(chatId int64, userId int64, level int, until_date int64) error
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
		logger.Errorf("SendMessage err: %v", err)
		return
	}

	if duration > 0 {
		// Commit delay delete task to timewheel.
		task := &delayDeleteTask{
			ChatId:     resultMsg.Chat.ID,
			MessageId:  resultMsg.MessageID,
			BotService: s,
		}
		timewheel.AddTask(duration, task)
	}
}

func (s *botService) DeleteMessageById(chatId int64, messageId int) {
	deleteReq := tgbotapi.NewDeleteMessage(chatId, messageId)

	s.sendDeleteReq(deleteReq)
}

func (s *botService) SetPermission(chatId int64, userId int64, level int, until_date int64) error {
	permission := getChatPermissionByLevel(level)
	operate := tgbotapi.RestrictChatMemberConfig{
		ChatMemberConfig: tgbotapi.ChatMemberConfig{ChatID: chatId, UserID: userId},
		UntilDate:        until_date,
		Permissions:      permission,
	}

	resp, err := s.BotAPI.Request(operate)
	if err != nil {
		logger.Errorf("SetPermission err: %s, resp: %v", err.Error(), resp)
		return err
	}

	return nil
}

func (s *botService) sendDeleteReq(deleteReq tgbotapi.Chattable) {
	retry.Do(func() error {
		// Send delete request to telegram.
		_, err := s.BotAPI.Request(deleteReq)
		if err != nil {
			if boterr.IsNotFound(err) || boterr.IsCantBeDelete(err) {
				return nil
			}
			logger.Errorf("DeleteMessage err: %s", err.Error())
			return err
		}
		return nil
	}, retry.Attempts(5), retry.Delay(time.Second))
}

func getChatPermissionByLevel(level int) *tgbotapi.ChatPermissions {
	result := &tgbotapi.ChatPermissions{}

	if level >= models.NONE {
		result.CanSendMessages = true
	}

	if level >= models.LIMIT {
		result.CanSendOtherMessages = true
	}

	if level >= models.JUNIOR {
		result.CanSendOtherMessages = true
		result.CanSendMediaMessages = true
		result.CanAddWebPagePreviews = true
	}

	return result
}

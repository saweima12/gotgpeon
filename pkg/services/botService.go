package services

import (
	"fmt"
	"gotgpeon/config"
	"gotgpeon/data/models"
	"gotgpeon/libs/timewheel"
	"gotgpeon/logger"
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
	SendMessage(message tgbotapi.Chattable, duration time.Duration) (msgId int)
	SendVbanMarkup(chatId int64, memberId int64, memberName string) (msgId int)
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

func (s *botService) SendMessage(message tgbotapi.Chattable, duration time.Duration) (msgId int) {
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
	return resultMsg.MessageID
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

func (s *botService) SendVbanMarkup(chatId int64, memberId int64, memberName string) (msgId int) {
	msg := getVbanMessage(chatId, memberId, memberName)
	resp, err := s.BotAPI.Send(msg)
	if err != nil {
		logger.Errorf("SetVbanMarkup err: %s", err.Error())
		return
	}

	return resp.MessageID
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

func getVbanMessage(chatId int64, memberId int64, memberName string) *tgbotapi.MessageConfig {
	tplStr := config.GetTextLang().ActVbanMarkup
	text := fmt.Sprintf(tplStr, chatId, memberId, memberName)

	markup := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Wee", "wwee"),
		),
	)

	return &tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID:      chatId,
			ReplyMarkup: markup,
		},
		ParseMode:             tgbotapi.ModeMarkdown,
		Text:                  text,
		DisableWebPagePreview: false,
	}
}

package utils

import (
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MessageHelper struct {
	*tgbotapi.Message
}

func NewMessageHelper(message *tgbotapi.Message) *MessageHelper {
	return &MessageHelper{
		Message: message,
	}
}

func (h *MessageHelper) IsSuperGroup() bool {
	return h.Chat.IsSuperGroup()
}

func (h *MessageHelper) FullName() string {
	return h.From.FirstName + " " + h.From.LastName
}

func (h *MessageHelper) UserId() string {
	userId := int(h.From.ID)
	return strconv.Itoa(userId)
}

func (h *MessageHelper) ChatId() string {
	chatId := int(h.Chat.ID)
	return strconv.Itoa(chatId)
}

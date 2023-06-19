package utils

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

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

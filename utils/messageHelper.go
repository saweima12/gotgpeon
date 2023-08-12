package utils

import (
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MessageHelper struct {
	*tgbotapi.Message
	BotAPI *tgbotapi.BotAPI
}

func NewMessageHelper(message *tgbotapi.Message, botAPI *tgbotapi.BotAPI) *MessageHelper {
	return &MessageHelper{
		Message: message,
		BotAPI:  botAPI,
	}
}

func (h *MessageHelper) IsSuperGroup() bool {
	return h.Chat.IsSuperGroup()
}

func (h *MessageHelper) IsForward() bool {
	return h.ForwardDate != 0
}

func (h *MessageHelper) FullName() string {
	return h.From.FirstName + " " + h.From.LastName
}

func (h *MessageHelper) UserId() int64 {
	return h.From.ID
}

func (h *MessageHelper) UserIdStr() string {
	userId := int(h.From.ID)
	return strconv.Itoa(userId)
}

func (h *MessageHelper) ChatId() int64 {
	return h.Chat.ID
}

func (h *MessageHelper) ChatIdStr() string {
	chatId := int(h.Chat.ID)
	return strconv.Itoa(chatId)
}

func (h *MessageHelper) ContentType() string {
	if h.ForwardDate != 0 {
		return "forward"
	}

	if h.Sticker != nil {
		return "sticker"
	}

	if h.Dice != nil {
		return "dice"
	}

	if h.Video != nil {
		return "video"
	}

	if h.Animation != nil {
		return "animation"
	}

	if h.Photo != nil {
		return "photo"
	}

	if h.Audio != nil {
		return "audio"
	}

	if h.Document != nil {
		return "document"
	}

	if h.Text != "" {
		return "text"
	}

	return "unknown"
}

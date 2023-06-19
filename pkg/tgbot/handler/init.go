package handler

import (
	"gotgpeon/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type MessageHandler interface {
	HandleMessage(message *tgbotapi.Message, bot *tgbotapi.BotAPI)
}

type messageHandler struct {
}

func NewMessageHandler(dbConn *gorm.DB, cacheConn *redis.Client) MessageHandler {
	// TODO: Initialize Repository and service.

	return &messageHandler{}
}

func (h *messageHandler) HandleMessage(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	helper := utils.NewMessageHelper(message)

	if helper.IsSuperGroup() {
		h.handleGroupMessage(helper)
	}
}

func (h *messageHandler) handleGroupMessage(helper *utils.MessageHelper) {

}

func (h *messageHandler) handleEnterGroupMsg(helper *utils.MessageHelper) {

}

func (h *messageHandler) handleLeaveGroupMsg(helper *utils.MessageHelper) {

}

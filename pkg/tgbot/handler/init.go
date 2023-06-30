package handler

import (
	"gotgpeon/pkg/repositories"
	"gotgpeon/pkg/services"
	"gotgpeon/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type MessageHandler interface {
	HandleMessage(message *tgbotapi.Message, bot *tgbotapi.BotAPI)
}

type messageHandler struct {
	peonService services.PeonService
}

func NewMessageHandler(dbConn *gorm.DB, cacheConn *redis.Client) MessageHandler {
	// Initialize Repositories
	chatRepo := repositories.NewChatRepo(dbConn, cacheConn)
	botRepo := repositories.NewBotConfigRepo(dbConn, cacheConn)

	// Initialize Services
	peonService := services.NewPeonService(chatRepo, botRepo)

	return &messageHandler{
		peonService: peonService,
	}
}

func (h *messageHandler) HandleMessage(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	helper := utils.NewMessageHelper(message)

	if helper.IsSuperGroup() {
		h.handleGroupMessage(helper)
	}
}

func (h *messageHandler) handleEnterGroupMsg(helper *utils.MessageHelper) {

}

func (h *messageHandler) handleLeaveGroupMsg(helper *utils.MessageHelper) {

}

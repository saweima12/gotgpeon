package handler

import (
	"gotgpeon/logger"
	"gotgpeon/models"
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
	peonService   services.PeonService
	recordService services.RecordService
}

func NewMessageHandler(dbConn *gorm.DB, cacheConn *redis.Client) MessageHandler {
	// Initialize Repositories
	chatRepo := repositories.NewChatRepo(dbConn, cacheConn)
	botRepo := repositories.NewBotConfigRepo(dbConn, cacheConn)
	recordRepo := repositories.NewRecordRepository(dbConn, cacheConn)

	// Initialize Services
	peonService := services.NewPeonService(chatRepo, botRepo)
	recordService := services.NewRecordService(recordRepo)

	return &messageHandler{
		peonService:   peonService,
		recordService: recordService,
	}
}

func (h *messageHandler) HandleMessage(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	helper := utils.NewMessageHelper(message)

	logger.Info(message)

	if helper.IsSuperGroup() {
		h.handleGroupMessage(helper)
	}
}

func (h *messageHandler) getMessageContext(helper *utils.MessageHelper, chatCfg *models.ChatConfig) *models.MessageContext {

	chatId := helper.ChatId()
	userId := helper.UserId()

	// Check userid in the allowList.
	isAllowlist := h.peonService.IsAllowListUser(userId)

	// Query UserRecord.
	recordQuery := &models.MessageRecord{
		UserId:   userId,
		FullName: helper.FullName(),
	}
	userRecord := h.recordService.GetUserRecord(chatId, recordQuery)

	return &models.MessageContext{
		ChatCfg:     chatCfg,
		IsWhitelist: isAllowlist,
		Record:      userRecord,
	}
}

func (h *messageHandler) handleEnterGroupMsg(helper *utils.MessageHelper) {

}

func (h *messageHandler) handleLeaveGroupMsg(helper *utils.MessageHelper) {

}

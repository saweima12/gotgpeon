package handler

import (
	"gotgpeon/logger"
	"gotgpeon/models"
	"gotgpeon/pkg/repositories"
	"gotgpeon/pkg/services"
	"gotgpeon/pkg/tgbot/command"
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
	cmdMap        *command.CommandMap
}

func NewMessageHandler(dbConn *gorm.DB, cacheConn *redis.Client) MessageHandler {
	// Initialize Repositories
	chatRepo := repositories.NewChatRepo(dbConn, cacheConn)
	botRepo := repositories.NewBotConfigRepo(dbConn, cacheConn)
	recordRepo := repositories.NewRecordRepository(dbConn, cacheConn)

	// Initialize Services
	peonService := services.NewPeonService(chatRepo, botRepo)
	recordService := services.NewRecordService(recordRepo)
	botService := services.NewBotService()

	cmdMap := &command.CommandMap{
		PeonService:   peonService,
		RecordService: recordService,
		BotService:    botService,
	}
	cmdMap.Init()

	// Initialize command map
	return &messageHandler{
		peonService:   peonService,
		recordService: recordService,
		cmdMap:        cmdMap,
	}
}

func (h *messageHandler) HandleMessage(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	helper := utils.NewMessageHelper(message, bot)

	logger.Debug(message)
	if helper.IsSuperGroup() {
		// Check if message is a command.
		if helper.IsCommand() {
			h.cmdMap.Invoke(helper)
			return
		}

		// message is not a command, check process.
		h.handleGroupMessage(helper)
	}
}

func (h *messageHandler) getMessageContext(helper *utils.MessageHelper, chatCfg *models.ChatConfig) *models.MessageContext {

	chatId := helper.ChatIdStr()
	userId := helper.UserIdStr()

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

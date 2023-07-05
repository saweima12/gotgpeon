package handler

import (
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
	commandMap    map[string]func(helper *utils.MessageHelper)
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

	result := &messageHandler{
		peonService:   peonService,
		recordService: recordService,
	}

	// Initialize command map
	result.commandMap = result.initGroupCommandMap()
	return result
}

func (h *messageHandler) HandleMessage(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	helper := utils.NewMessageHelper(message, bot)

	if helper.IsSuperGroup() {
		// Check if message is a command.
		if helper.IsCommand() {
			h.handleGroupCommand(helper)
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

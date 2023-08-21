package handler

import (
	"gotgpeon/config"
	"gotgpeon/logger"
	"gotgpeon/models"
	"gotgpeon/pkg/repositories"
	"gotgpeon/pkg/services"
	"gotgpeon/pkg/tgbot/checker"
	"gotgpeon/pkg/tgbot/command"
	"gotgpeon/utils"
	"gotgpeon/utils/jsonutil"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type MessageHandler interface {
	HandleMessage(message *tgbotapi.Message, bot *tgbotapi.BotAPI, isEdit bool)
}

type messageHandler struct {
	peonService    services.PeonService
	recordService  services.RecordService
	botService     services.BotService
	deletedService services.DeletedService
	checker        checker.CheckerHandler
	cmdMap         command.CommandHandler
}

func NewMessageHandler(dbConn *gorm.DB, cacheConn *redis.Client, botAPI *tgbotapi.BotAPI) MessageHandler {
	// Initialize Repositories
	chatRepo := repositories.NewChatRepo(dbConn, cacheConn)
	botRepo := repositories.NewBotConfigRepo(dbConn, cacheConn)
	recordRepo := repositories.NewRecordRepository(dbConn, cacheConn)
	deletedMsgRepo := repositories.NewDeletedMsgRepository(dbConn, cacheConn)

	// Initialize Services
	peonService := services.NewPeonService(chatRepo, botRepo, deletedMsgRepo)
	recordService := services.NewRecordService(recordRepo)
	botService := services.NewBotService(botAPI)
	deletedService := services.NewDeletedService(deletedMsgRepo)

	// Initialize commandMap
	cmdMap := &command.CommandMap{
		PeonService:   peonService,
		RecordService: recordService,
		BotService:    botService,
	}
	cmdMap.Init()

	// Initialize checker.
	checker := &checker.MessageChecker{}
	checker.Init()

	// Initialize command map
	return &messageHandler{
		peonService:    peonService,
		recordService:  recordService,
		botService:     botService,
		deletedService: deletedService,
		cmdMap:         cmdMap,
		checker:        checker,
	}
}

func (h *messageHandler) HandleMessage(message *tgbotapi.Message, bot *tgbotapi.BotAPI, isEdit bool) {
	helper := utils.NewMessageHelper(message, bot)

	// HandleMessage.
	data, _ := jsonutil.MarshalToString(message)
	logger.Debug(data)

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
	chatId := helper.ChatId()
	userId := helper.UserId()

	// Check userid in the allowList.
	isAllowlist := h.peonService.IsAllowListUser(userId)

	// Query UserRecord.
	recordQuery := &models.MessageRecord{
		MemberId: userId,
		FullName: helper.FullName(),
	}
	userRecord := h.recordService.GetUserRecord(chatId, recordQuery)

	// Get serviceConfig.
	commonCfg := config.GetConfig().Common

	return &models.MessageContext{
		ChatCfg:     chatCfg,
		CommonCfg:   &commonCfg,
		IsWhitelist: isAllowlist,
		Record:      userRecord,
	}
}

func (h *messageHandler) handleEnterGroupMsg(helper *utils.MessageHelper) {

}

func (h *messageHandler) handleLeaveGroupMsg(helper *utils.MessageHelper) {

}

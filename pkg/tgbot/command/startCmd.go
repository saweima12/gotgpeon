package command

import (
	"fmt"
	"gotgpeon/config"
	"gotgpeon/logger"
	"gotgpeon/models"
	"gotgpeon/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Handle /start command
func (h *CommandMap) handleStartCmd(helper *utils.MessageHelper) {
	// Check user is allowlist user..
	if !h.PeonService.IsAllowListUser(helper.UserId()) {
		return
	}

	// Get chat adminstrator from botapi
	chatAdminstrator, err := helper.BotAPI.GetChatAdministrators(tgbotapi.ChatAdministratorsConfig{
		ChatConfig: tgbotapi.ChatConfig{ChatID: helper.ChatId()},
	})

	if err != nil {
		logger.Errorf("Register ChatGroup error: %s", err.Error())
		return
	}

	strAdminstratorIds := make([]int64, 0, len(chatAdminstrator))
	// process user id
	for _, user := range chatAdminstrator {
		userId := user.User.ID
		strAdminstratorIds = append(strAdminstratorIds, userId)
	}
	// save to database and cache
	chatId := helper.ChatId()
	chatCfg := h.PeonService.GetChatConfig(chatId, helper.Chat.Title)
	chatCfg.Status = models.OK
	chatCfg.Adminstrators = strAdminstratorIds

	err = h.PeonService.SetChatConfig(chatCfg)
	if err != nil {
		logger.Errorf("Register ChatGroup SetConfig error: %s", err.Error())
		return
	}

	// Send tips message.
	sendText := fmt.Sprintf(config.GetTextLang().TipsStartCmd, helper.Chat.Title)
	newMsg := tgbotapi.NewMessage(helper.ChatId(), sendText)
	newMsg.ParseMode = tgbotapi.ModeHTML
	h.BotService.SendMessage(newMsg, 0)
}

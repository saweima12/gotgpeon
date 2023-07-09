package command

import (
	"gotgpeon/models"
	"gotgpeon/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Handle /start command
func (h *CommandMap) handleStartCommand(helper *utils.MessageHelper) {
	// Check user is allowlist user..
	if !h.PeonService.IsAllowListUser(helper.UserIdStr()) {
		return
	}

	chatIdStr := helper.ChatIdStr()
	chatCfg := h.PeonService.GetChatConfig(chatIdStr, helper.Chat.Title)
	chatCfg.Status = models.OK

	h.PeonService.SetChatConfig(chatCfg)
	// TODO: Send tips message.
	newMsg := tgbotapi.NewMessage(helper.ChatId(), "Weeed")
	helper.BotAPI.Send(newMsg)
}

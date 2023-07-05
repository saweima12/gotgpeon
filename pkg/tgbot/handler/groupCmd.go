package handler

import (
	"fmt"
	"gotgpeon/models"
	"gotgpeon/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *messageHandler) initGroupCommandMap() map[string]func(helper *utils.MessageHelper) {
	return map[string]func(helper *utils.MessageHelper){
		"start":    h.handleStartCommand,
		"point":    h.handlePointCommand,
		"setlevel": h.handleSetLevelCommand,
		"save":     h.handleSaveCommand,
	}
}

func (h *messageHandler) handleGroupCommand(helper *utils.MessageHelper) {
	cmdStr := helper.Command()
	cmdFunc, ok := h.commandMap[cmdStr]
	if ok {
		cmdFunc(helper)
	}
}

// Handle /start command
func (h *messageHandler) handleStartCommand(helper *utils.MessageHelper) {
	// Check user is allowlist user..
	if !h.peonService.IsAllowListUser(helper.UserIdStr()) {
		return
	}

	chatIdStr := helper.ChatIdStr()
	chatCfg := h.peonService.GetChatConfig(chatIdStr, helper.Chat.Title)
	chatCfg.Status = models.OK

	h.peonService.SetChatConfig(chatCfg)

	// TODO: Send tips message.
	newMsg := tgbotapi.NewMessage(helper.ChatId(), "Weeed")
	helper.BotAPI.Send(newMsg)
}

// handle /point command
func (h *messageHandler) handlePointCommand(helper *utils.MessageHelper) {
	fmt.Println("On point Command")
}

// handle /setlevel command
func (h *messageHandler) handleSetLevelCommand(helper *utils.MessageHelper) {
	fmt.Println("On Start Command")
}

// handle /save command
func (h *messageHandler) handleSaveCommand(helper *utils.MessageHelper) {
	fmt.Println("On Save command.")
}

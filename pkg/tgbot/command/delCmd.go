package command

import (
	"gotgpeon/utils"
)

// handle /point command
func (h *CommandMap) handleDelCmd(helper *utils.MessageHelper) {

	chatCfg := h.PeonService.GetChatConfig(helper.ChatId(), helper.Chat.Title)
	if !chatCfg.IsAvaliable() {
		return
	}

	userId := helper.UserId()
	isAllowListUser := h.PeonService.IsAllowListUser(userId)
	isAdminstrator := chatCfg.IsAdminstrator(userId)

	if !isAllowListUser && !isAdminstrator {
		return
	}

	// Define parameter
	if helper.ReplyToMessage != nil {
		h.BotService.DeleteMessageById(helper.ChatId(), helper.ReplyToMessage.MessageID)
		h.BotService.DeleteMessageById(helper.ChatId(), helper.MessageID)
	}
}

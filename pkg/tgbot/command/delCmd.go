package command

import (
	"gotgpeon/utils"
)

// handle /point command
func (h *CommandMap) handleDelCmd(helper *utils.MessageHelper) {

	chatCfg := h.PeonService.GetChatConfig(helper.ChatIdStr(), helper.Chat.Title)
	if !chatCfg.IsAvaliable() {
		return
	}

	userIdStr := helper.UserIdStr()
	isAllowListUser := h.PeonService.IsAllowListUser(helper.UserIdStr())
	isAdminstrator := chatCfg.IsAdminstrator(userIdStr)

	if !isAllowListUser && !isAdminstrator {
		return
	}

	// Define parameter
	if helper.ReplyToMessage != nil {
		h.BotService.DeleteMessageById(helper.ChatId(), helper.ReplyToMessage.MessageID)
		h.BotService.DeleteMessageById(helper.ChatId(), helper.MessageID)
	}
}

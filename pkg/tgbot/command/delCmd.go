package command

import "gotgpeon/pkg/tgbot/core"

// handle /point command
func (h *CommandMap) handleDelCmd(helper *core.MessageHelper) {

	chatCfg := h.PeonService.GetChatConfig(helper.ChatId())
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

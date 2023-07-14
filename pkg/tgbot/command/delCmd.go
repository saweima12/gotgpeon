package command

import (
	"gotgpeon/utils"
)

// handle /point command
func (h *CommandMap) handleDelCommand(helper *utils.MessageHelper) {
	// Define parameter
	if helper.ReplyToMessage != nil {
		h.BotService.DeleteMessage(helper.ChatId(), helper.ReplyToMessage.MessageID)
	}
}

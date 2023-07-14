package handler

import (
	"gotgpeon/logger"
	"gotgpeon/utils"
)

func (h *messageHandler) handleGroupMessage(helper *utils.MessageHelper) {

	chatId := helper.ChatIdStr()
	// Check chat is avaliable
	chatCfg := h.peonService.GetChatConfig(chatId, helper.Chat.Title)

	// if chatCfg.Status != models.OK {
	// 	return
	// }
	ctx := h.getMessageContext(helper, chatCfg)
	result := h.checker.CheckMessage(helper, ctx)

	if result.MarkDelete {

	}

	if !result.MarkRecord {
		return
	}
	// Add point.
	err := h.recordService.AddUserPoint(chatId, ctx.Record)
	if err != nil {
		logger.Errorf("HandleGroupMessage Err: %s", err.Error())
	}
}

package handler

import (
	"gotgpeon/logger"
	"gotgpeon/models"
	"gotgpeon/utils"
)

func (h *messageHandler) handleGroupMessage(helper *utils.MessageHelper) {

	chatId := helper.ChatIdStr()
	// Check chat is avaliable
	chatCfg := h.peonService.GetChatConfig(chatId, helper.Chat.Title)

	// check group is avaliable.
	if chatCfg.Status != models.OK {
		return
	}

	ctx := h.getMessageContext(helper, chatCfg)
	result := h.checker.CheckMessage(helper, ctx)

	if result.MarkDelete {
		h.botService.DeleteMessageById(helper.ChatId(), helper.MessageID)
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

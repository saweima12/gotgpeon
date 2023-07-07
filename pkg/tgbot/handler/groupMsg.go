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

	// if chatCfg.Status != models.OK && !isAllowUser {
	// 	return
	// }

	ctx := h.getMessageContext(helper, chatCfg)

	// TODO: Check message data.

	// Add point.
	err := h.recordService.AddUserPoint(chatId, ctx.Record)
	if err != nil {
		logger.Errorf("HandleGroupMessage Err: %s", err.Error())
	}
}

func (h *messageHandler) checkGroupMessage(helper *utils.MessageHelper, ctx *models.MessageContext) {

}

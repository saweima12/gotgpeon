package handler

import (
	"gotgpeon/models"
	"gotgpeon/utils"
	"strconv"
)

func (h *messageHandler) handleGroupMessage(helper *utils.MessageHelper) {

}

func (h *messageHandler) getMessageContext(helper *utils.MessageHelper) {

	chatId := strconv.Itoa(int(helper.Chat.ID))
	chatName := helper.Chat.Title

	userId := strconv.Itoa(int(helper.From.ID))

	// Get data
	chatCfg := h.peonService.GetChatConfig(chatId, chatName)

	return models.MessageContext{
		ChatCfg: chatCfg,
	}
}

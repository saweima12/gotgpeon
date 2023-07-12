package handler

import (
	"gotgpeon/logger"
	"gotgpeon/models"
	"gotgpeon/pkg/tgbot/checker"
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
	result := &checker.CheckResult{
		MustDelete: false,
		MustRecord: true,
	}
	// Check message is avaliable.
	if ctx.Record.MemberLevel <= models.LIMIT && !ctx.IsAdminstrator() {
		result = h.checker.CheckMessage(helper, ctx)
	}

	if !result.MustRecord {
		return
	}
	// Add point.
	err := h.recordService.AddUserPoint(chatId, ctx.Record)
	if err != nil {
		logger.Errorf("HandleGroupMessage Err: %s", err.Error())
	}
}

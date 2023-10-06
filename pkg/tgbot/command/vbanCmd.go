package command

import (
	"gotgpeon/data/models"
	"gotgpeon/pkg/tgbot/core"
)

func (m *CommandMap) handVbanCmd(helper *core.MessageHelper) {
	// Check Chat isavaliable
	chatId := helper.ChatId()
	chatCfg := m.PeonService.GetChatConfig(chatId)
	if !chatCfg.IsAvaliable() {
		return
	}

	// Check sender permission
	userId := helper.UserId()
	query := &models.MessageRecord{
		MemberId: userId,
	}
	record := m.RecordService.GetUserRecordByCaht(chatId, query)
	isAdminstrator := chatCfg.IsAdminstrator(userId)
	isSenior := record.MemberLevel < models.SENIOR
	if !isAdminstrator && isSenior {
		return
	}

	// Execute Vban command

}

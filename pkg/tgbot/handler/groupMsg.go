package handler

import (
	"fmt"
	"gotgpeon/config"
	"gotgpeon/libs/ants"
	"gotgpeon/libs/json"
	"gotgpeon/logger"
	"gotgpeon/models"
	"gotgpeon/pkg/tgbot/core"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *messageHandler) handleGroupMessage(helper *core.MessageHelper) {

	chatId := helper.ChatId()
	// Check chat is avaliable
	chatCfg := h.peonService.GetChatConfig(chatId)

	// check group is avaliable.
	if chatCfg.Status != models.OK {
		return
	}

	ctx := h.getMessageContext(helper, chatCfg)
	result := h.checker.CheckMessage(helper, ctx)

	if result.MarkDelete {
		h.handleDeleteMessage(helper, result.Message)
	}

	if result.MarkRecord {
		ctx.Record.Point += 1
	}

	// Update user record.
	err := h.recordService.SetUserRecordCache(chatId, ctx.Record)
	if err != nil {
		logger.Errorf("HandleGroupMessage Err: %s", err.Error())
	}
}

func (h *messageHandler) handleDeleteMessage(helper *core.MessageHelper, text string) {
	ants.Submit(func() {
		// print log
		jsonStr, err := json.MarshalToString(helper.Message)
		if err != nil {
			logger.Errorf("handleDeleteMessage err: %v", helper)
		}
		logger.Infof("DeleteMessage: %s", jsonStr)

		// Insert into database
		err = h.deletedService.InsertDeletedRecord(helper.ChatId(), helper.ContentType(), helper.Message)
		if err != nil {
			logger.Errorf("handleDeleteMessage insertIntoDb err: %s", err.Error())
		}

		// send delete request.
		h.botService.DeleteMessageById(helper.ChatId(), helper.MessageID)

		msgText := fmt.Sprintf(config.GetTextLang().TipsDeleteMsg,
			helper.FullName(),
			helper.UserIdStr(),
			text,
		)

		// delete finish, send tips.
		newMsg := tgbotapi.NewMessage(helper.ChatId(), msgText)
		newMsg.ParseMode = tgbotapi.ModeMarkdown
		h.botService.SendMessage(newMsg, time.Second*5)
	})
}

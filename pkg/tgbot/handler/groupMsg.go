package handler

import (
	"fmt"
	"gotgpeon/config"
	"gotgpeon/logger"
	"gotgpeon/models"
	"gotgpeon/utils"
	"gotgpeon/utils/jsonutil"
	"gotgpeon/utils/poolutil"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *messageHandler) handleGroupMessage(helper *utils.MessageHelper) {

	chatId := helper.ChatId()
	// Check chat is avaliable
	chatCfg := h.peonService.GetChatConfig(chatId, helper.Chat.Title)

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

func (h *messageHandler) handleDeleteMessage(helper *utils.MessageHelper, text string) {
	poolutil.Submit(func() {
		// print log
		jsonStr, err := jsonutil.MarshalToString(helper.Message)
		if err != nil {
			logger.Errorf("handleDeleteMessage err: %v", helper)
		}
		logger.Infof("DeleteMessage: %s", jsonStr)

		// Insert into database
		err = h.peonService.InsertDeletedRecord(helper.ChatId(), helper.ContentType(), helper.Message)
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

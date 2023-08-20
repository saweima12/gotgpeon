package command

import (
	"gotgpeon/models"
	"gotgpeon/utils"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// handle /point command
func (h *CommandMap) handlePointCmd(helper *utils.MessageHelper) {

	// Define parameter
	chatId := helper.ChatId()
	chatCfg := h.PeonService.GetChatConfig(chatId)

	// Check group is avaliable.
	if !chatCfg.IsAvaliable() {
		return
	}

	query := &models.MessageRecord{
		MemberId: helper.UserId(),
		FullName: helper.FullName(),
	}
	// Create tips message.
	userRecord := h.RecordService.GetUserRecord(chatId, query)

	text := "Point: " + strconv.Itoa(userRecord.Point)
	newMsg := tgbotapi.NewMessage(helper.ChatId(), text)

	// Send tips
	h.BotService.SendMessage(newMsg, time.Second*5)
}

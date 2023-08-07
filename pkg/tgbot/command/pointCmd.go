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
	chatIdStr := helper.ChatIdStr()
	chatCfg := h.PeonService.GetChatConfig(chatIdStr, helper.Chat.Title)

	// Check group is avaliable.
	if !chatCfg.IsAvaliable() {
		return
	}

	query := &models.MessageRecord{
		UserId:   helper.UserIdStr(),
		FullName: helper.FullName(),
	}
	// Create tips message.
	userRecord := h.RecordService.GetUserRecord(chatIdStr, query)

	text := "Point: " + strconv.Itoa(userRecord.Point)
	newMsg := tgbotapi.NewMessage(helper.ChatId(), text)

	// Send tips
	h.BotService.SendMessage(newMsg, time.Second*5)
}

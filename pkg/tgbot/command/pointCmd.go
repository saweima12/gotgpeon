package command

import (
	"gotgpeon/models"
	"gotgpeon/utils"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// handle /point command
func (h *CommandMap) handlePointCommand(helper *utils.MessageHelper) {
	// Define parameter
	chatIdStr := helper.ChatIdStr()
	query := &models.MessageRecord{
		UserId: helper.UserIdStr(),
	}

	userRecord := h.RecordService.GetUserRecord(chatIdStr, query)
	text := "Point: " + strconv.Itoa(userRecord.Point)
	// Send tips
	newMsg := tgbotapi.NewMessage(helper.ChatId(), text)
	helper.BotAPI.Send(newMsg)
}

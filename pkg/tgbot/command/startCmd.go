package command

import (
	"fmt"
	"gotgpeon/config"
	"gotgpeon/logger"
	"gotgpeon/models"
	"gotgpeon/utils"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Handle /start command
func (h *CommandMap) handleStartCommand(helper *utils.MessageHelper) {
	// Check user is allowlist user..
	if !h.PeonService.IsAllowListUser(helper.UserIdStr()) {
		return
	}

	// Get chat adminstrator from botapi
	chatAdminstrator, err := helper.BotAPI.GetChatAdministrators(tgbotapi.ChatAdministratorsConfig{
		ChatConfig: tgbotapi.ChatConfig{ChatID: helper.ChatId()},
	})

	if err != nil {
		logger.Errorf("Register ChatGroup error: %s", err.Error())
		return
	}

	strAdminstratorIds := make([]string, 0, len(chatAdminstrator))
	// process user id
	for index, user := range chatAdminstrator {
		fmt.Println(index)
		userId := strconv.Itoa(int(user.User.ID))
		if strings.Trim(userId, " ") == "" {
			continue
		}
		strAdminstratorIds = append(strAdminstratorIds, userId)
	}
	// save to database and cache
	chatIdStr := helper.ChatIdStr()
	chatCfg := h.PeonService.GetChatConfig(chatIdStr, helper.Chat.Title)
	chatCfg.Status = models.OK
	chatCfg.Adminstrators = strAdminstratorIds

	err = h.PeonService.SetChatConfig(chatCfg)
	if err != nil {
		logger.Errorf("Register ChatGroup SetConfig error: %s", err.Error())
		return
	}

	// Send tips message.
	sendText := fmt.Sprintf(config.GetTextLang().TipsStartCmd, helper.Chat.Title)
	newMsg := tgbotapi.NewMessage(helper.ChatId(), sendText)
	newMsg.ParseMode = tgbotapi.ModeHTML
	helper.BotAPI.Send(newMsg)
}

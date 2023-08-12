package command

import (
	"fmt"
	"gotgpeon/config"
	"gotgpeon/logger"
	"gotgpeon/models"
	"gotgpeon/utils"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *CommandMap) handleSetLevelCmd(helper *utils.MessageHelper) {
	chatIdStr := helper.ChatIdStr()
	chatCfg := h.PeonService.GetChatConfig(chatIdStr, helper.Chat.Title)

	// Check group is avaliable.
	if !chatCfg.IsAvaliable() {
		return
	}

	// Check parameter is ok.
	level, isOK := validateSetLevelParameter(helper)
	if !isOK {
		return
	}

	// Check sender permission.
	userIdStr := helper.UserIdStr()
	isAllowListUser := h.PeonService.IsAllowListUser(userIdStr)
	isAdminstrator := chatCfg.IsAdminstrator(userIdStr)
	if !isAdminstrator && !isAllowListUser {
		return
	}

	targetIdStr := strconv.Itoa(int(helper.ReplyToMessage.From.ID))
	replyMsg := helper.ReplyToMessage
	targetName := replyMsg.From.FirstName + " " + replyMsg.From.LastName
	query := models.MessageRecord{
		UserId:   targetIdStr,
		FullName: targetName,
	}
	// set memberlevel & save
	userRecord := h.RecordService.GetUserRecord(chatIdStr, &query)
	userRecord.MemberLevel = level
	userRecord.FullName = targetName

	err := h.RecordService.SetUserRecordCache(chatIdStr, userRecord)
	if err != nil {
		logger.Errorf("Set member %s level err: %s", targetName, err.Error())
		return
	}

	err = h.RecordService.SetUserRecordDB(chatIdStr, userRecord)
	if err != nil {
		logger.Errorf("Set member %s level err: %s", targetName, err.Error())
		return
	}

	err = h.BotService.SetPermission(helper.ChatId(), replyMsg.From.ID, level, 0)
	if err != nil {
		logger.Errorf("Set member %s level err: %s", targetName, err.Error())
		return
	}

	levelText := strings.ToUpper(helper.CommandArguments())
	// Send success tips.
	msgText := fmt.Sprintf(
		config.GetTextLang().TipsSetPermissionCmd,
		targetName,
		targetIdStr,
		levelText,
	)
	newMsg := tgbotapi.NewMessage(helper.ChatId(), msgText)
	newMsg.ParseMode = tgbotapi.ModeMarkdown
	h.BotService.SendMessage(newMsg, time.Second*5)
}

func validateSetLevelParameter(helper *utils.MessageHelper) (level int, ok bool) {

	if helper.ReplyToMessage == nil {
		return models.NONE, false
	}

	arguments := helper.CommandArguments()
	level, ok = models.MemberLevelMap[arguments]
	if !ok {
		return models.NONE, false
	}

	return level, true
}

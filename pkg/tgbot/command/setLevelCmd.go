package command

import (
	"fmt"
	"gotgpeon/config"
	"gotgpeon/logger"
	"gotgpeon/models"
	"gotgpeon/utils"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *CommandMap) handleSetLevelCmd(helper *utils.MessageHelper) {
	chatId := helper.ChatId()
	chatCfg := h.PeonService.GetChatConfig(chatId)

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
	userId := helper.UserId()
	isAllowListUser := h.PeonService.IsAllowListUser(userId)
	isAdminstrator := chatCfg.IsAdminstrator(userId)
	if !isAdminstrator && !isAllowListUser {
		return
	}

	replyMsg := helper.ReplyToMessage
	targetName := replyMsg.From.FirstName + " " + replyMsg.From.LastName
	query := models.MessageRecord{
		MemberId: replyMsg.From.ID,
		FullName: targetName,
	}
	// set memberlevel & save
	userRecord := h.RecordService.GetUserRecord(chatId, &query)
	userRecord.MemberLevel = level
	userRecord.FullName = targetName

	err := h.RecordService.SetUserRecordCache(chatId, userRecord)
	if err != nil {
		logger.Errorf("Set member %s level err: %s", targetName, err.Error())
		return
	}

	err = h.RecordService.SetUserRecordDB(chatId, userRecord)
	if err != nil {
		logger.Errorf("Set member %s level err: %s", targetName, err.Error())
		return
	}

	targetId := replyMsg.From.ID
	err = h.BotService.SetPermission(helper.ChatId(), targetId, level, 0)
	if err != nil {
		logger.Errorf("Set member %s level err: %s", targetName, err.Error())
		return
	}

	levelText := strings.ToUpper(helper.CommandArguments())
	// Send success tips.
	msgText := fmt.Sprintf(
		config.GetTextLang().TipsSetPermissionCmd,
		targetName,
		targetId,
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

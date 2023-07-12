package checker

import (
	"gotgpeon/config"
	"gotgpeon/models"
	"gotgpeon/utils"
	"gotgpeon/utils/sliceutil"
	"strconv"
)

func (c *MessageChecker) CheckViabotOK(helper *utils.MessageHelper, ctx *models.MessageContext, parameter any) bool {
	if helper.ViaBot != nil {
		botId := strconv.Itoa(int(helper.ViaBot.ID))
		allowViaList := config.GetConfig().Common.AllowViaIds
		if sliceutil.Contains(botId, allowViaList) {
			return true
		}
		return false
	}
	return true
}

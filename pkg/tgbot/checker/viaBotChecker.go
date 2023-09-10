package checker

import (
	"gotgpeon/config"
	"gotgpeon/data/models"
	"gotgpeon/pkg/tgbot/core"
	"gotgpeon/utils/sliceutil"
)

func (c *MessageChecker) CheckViabotOK(helper *core.MessageHelper, ctx *models.MessageContext, result *CheckResult, parameter any) bool {
	if helper.ViaBot != nil {
		botUsername := helper.ViaBot.UserName
		allowViaList := config.GetConfig().Common.AllowViaBots
		if !sliceutil.Contains(botUsername, allowViaList) {
			result.MarkDelete = true
		}
	}
	return true
}

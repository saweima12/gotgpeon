package checker

import (
	"gotgpeon/config"
	"gotgpeon/models"
	"gotgpeon/utils"
	"gotgpeon/utils/sliceutil"
)

func (c *MessageChecker) CheckViabotOK(helper *utils.MessageHelper, ctx *models.MessageContext, result *CheckResult, parameter any) bool {
	if helper.ViaBot != nil {
		botUsername := helper.ViaBot.UserName
		allowViaList := config.GetConfig().Common.AllowViaBots
		if !sliceutil.Contains(botUsername, allowViaList) {
			result.MarkDelete = true
		}
	}
	return true
}

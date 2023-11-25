package checker

import (
	"encoding/json"
	"gotgpeon/config"
	"gotgpeon/data/models"
	"gotgpeon/utils/sliceutil"
)

func (c *MessageChecker) CheckViabotOK(ctx *models.MessageContext, result *CheckResult, parameter json.RawMessage) bool {
	if ctx.Message.ViaBot != nil {
		botUsername := ctx.Message.ViaBot.UserName
		allowViaList := config.GetConfig().Common.AllowViaBots
		if !sliceutil.Contains(botUsername, allowViaList) {
			result.MarkDelete = true
		}
	}
	return true
}

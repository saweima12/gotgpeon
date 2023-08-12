package checker

import (
	"gotgpeon/config"
	"gotgpeon/models"
	"gotgpeon/utils"
	"gotgpeon/utils/sliceutil"
)

func (c *MessageChecker) CheckNoForward(helper *utils.MessageHelper, ctx *models.MessageContext, result *CheckResult, parameter any) bool {
	if helper.IsForward() {
		result.MarkDelete = true
		result.Message = config.GetTextLang().ErrForward
		return false
	}

	return true
}

// Check the message for any issues and return whether to continue the inspection.
func (c *MessageChecker) CheckTypeNoMedia(helper *utils.MessageHelper, ctx *models.MessageContext, result *CheckResult, parameter any) bool {
	// check message type
	if ctx.Record.MemberLevel >= models.LIMIT {
		if c.checkLimitUserOK(helper, ctx) {
			result.MarkDelete = false
			return false
		}
	}

	// Check type is not text.
	if helper.Text == "" {
		result.MarkDelete = true
		result.Message = config.GetTextLang().ErrTypeMedia
		return false
	}

	return true
}

func (c *MessageChecker) checkLimitUserOK(helper *utils.MessageHelper, ctx *models.MessageContext) bool {
	if helper.Sticker != nil || helper.Photo != nil || helper.Animation != nil || helper.Video != nil {
		if helper.ViaBot != nil {
			allowViaList := config.GetConfig().Common.AllowViaBots
			return sliceutil.Contains(helper.ViaBot.UserName, allowViaList)
		}
		return false
	}
	return false
}

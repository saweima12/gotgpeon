package checker

import (
	"encoding/json"
	"gotgpeon/config"
	"gotgpeon/data/models"
	"gotgpeon/pkg/tgbot/core"
	"gotgpeon/utils/sliceutil"
)

func (c *MessageChecker) CheckNoForward(ctx *models.MessageContext, result *CheckResult, parameter json.RawMessage) bool {
	if ctx.Message.IsForward() {
		result.MarkDelete = true
		result.Message = config.GetTextLang().ErrForward
		return false
	}

	return true
}

// Check the message for any issues and return whether to continue the inspection.
func (c *MessageChecker) CheckTypeNoMedia(ctx *models.MessageContext, result *CheckResult, parameter json.RawMessage) bool {
	// check message type
	helper := ctx.Message

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

func (c *MessageChecker) checkLimitUserOK(helper *core.MessageHelper, ctx *models.MessageContext) bool {
	if helper.Sticker != nil || helper.Photo != nil || helper.Animation != nil || helper.Video != nil {
		if helper.ViaBot != nil {
			allowViaList := config.GetConfig().Common.AllowViaBots
			return sliceutil.Contains(helper.ViaBot.UserName, allowViaList)
		}
		return false
	}
	return false
}

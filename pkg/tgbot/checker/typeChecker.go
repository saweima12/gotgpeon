package checker

import (
	"gotgpeon/config"
	"gotgpeon/models"
	"gotgpeon/utils"
	"gotgpeon/utils/sliceutil"
	"strconv"
)

func (c *MessageChecker) CheckTypeOK(helper *utils.MessageHelper, ctx *models.MessageContext, parameter any) bool {
	// check isforward ?
	if helper.IsForward() {
		return false
	}

	// check message type
	if ctx.Record.MemberLevel >= models.LIMIT {
		if c.checkLimitUserOK(helper, ctx) {
			return true
		}
	}

	// Check type is text.
	if helper.Text != "" {
		return true
	}

	return false
}

func (c *MessageChecker) checkLimitUserOK(helper *utils.MessageHelper, ctx *models.MessageContext) bool {
	if helper.Sticker != nil || helper.Photo != nil || helper.Animation != nil || helper.Video != nil {
		if helper.ViaBot != nil {
			allowViaList := config.GetConfig().Common.AllowViaIds
			viaBotIdStr := strconv.Itoa(int(helper.ViaBot.ID))
			return sliceutil.Contains(viaBotIdStr, allowViaList)
		}
		return false
	}
	return false
}

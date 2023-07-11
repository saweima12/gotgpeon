package checker

import (
	"fmt"
	"gotgpeon/config"
	"gotgpeon/logger"
	"gotgpeon/models"
	"gotgpeon/utils"
	"gotgpeon/utils/sliceutil"
	"strconv"
)

type CheckTypeParams struct {
	Data string `json:"data"`
}

func (c *MessageChecker) CheckTypeOK(helper *utils.MessageHelper, ctx *models.MessageContext, parameter any) bool {
	// process parameter.
	var params CheckTypeParams
	err := utils.AnyToStruct(parameter, &params)
	if err != nil {
		logger.Error(err)
		return true
	}

	// check isforward ?
	if helper.IsForward() {
		return false
	}

	// check message type
	if ctx.Record.MemberLevel >= models.LIMIT {
		if c.checkLimitUserOK(helper, ctx, &params) {
			return true
		}
	}

	if helper.Text != "" {
		fmt.Println("IsOkMessage:" + helper.Text)
		return true
	}

	return false

}

func (c *MessageChecker) checkLimitUserOK(helper *utils.MessageHelper, ctx *models.MessageContext, parameter *CheckTypeParams) bool {
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

package checker

import (
	"gotgpeon/models"
	"gotgpeon/utils"
	"gotgpeon/utils/goccutil"
)

func (c *MessageChecker) CheckContentOK(helper *utils.MessageHelper, ctx *models.MessageContext, parameter any) bool {
	if helper.Text != "" {
		originStr := helper.Text
		conStr := goccutil.S2T(originStr)

	}

	return true
}

func (c *MessageChecker) CheckNameOK(helper *utils.MessageHelper, ctx *models.MessageContext, parameter any) bool {
	return true
}

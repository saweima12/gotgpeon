package checker

import (
	"gotgpeon/config"
	"gotgpeon/models"
	"gotgpeon/utils"
)

func (c *MessageChecker) CheckEntitiesOK(helper *utils.MessageHelper, ctx *models.MessageContext, result *CheckResult, parameter any) bool {
	for _, entity := range helper.Entities {
		if entity.IsURL() || entity.IsMention() || entity.IsTextLink() || entity.IsHashtag() {
			result.MarkDelete = true
			result.Message = config.GetTextLang().TipsSetPermissionCmd
			return false
		}
	}
	return true
}

package checker

import (
	"encoding/json"
	"gotgpeon/config"
	"gotgpeon/data/models"
)

func (c *MessageChecker) CheckEntitiesOK(ctx *models.MessageContext, result *CheckResult, parameter json.RawMessage) bool {
	for _, entity := range ctx.Message.Entities {
		if entity.IsURL() || entity.IsMention() || entity.IsTextLink() || entity.IsHashtag() {
			result.MarkDelete = true
			result.Message = config.GetTextLang().TipsSetPermissionCmd
			return false
		}
	}
	return true
}

package checker

import (
	"gotgpeon/models"
	"gotgpeon/utils"
)

func (c *MessageChecker) CheckEntitiesOK(helper *utils.MessageHelper, ctx *models.MessageContext, parameter any) bool {
	for _, entity := range helper.Entities {
		if entity.IsURL() || entity.IsMention() || entity.IsTextLink() || entity.IsHashtag() {
			return false
		}
	}
	return true
}

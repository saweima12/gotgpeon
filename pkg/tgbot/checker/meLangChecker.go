package checker

import (
	"gotgpeon/config"
	"gotgpeon/data/models"
	"gotgpeon/pkg/tgbot/core"
	"regexp"
)

var MePtn = regexp.MustCompile("[\u0600-\u06FF\u0750-\u077F\uFB50-\uFDFF]")

func (c *MessageChecker) CheckContentNoMelang(helper *core.MessageHelper, ctx *models.MessageContext, result *CheckResult, parameter any) bool {
	if helper.Text != "" {
		// check middle east language.
		if MePtn.MatchString(helper.Text) {
			result.MarkDelete = true
			result.Message = config.GetTextLang().ErrContentNozhtw
			return false
		}
	}
	return true
}

func (c *MessageChecker) CheckNameNoMelang(helper *core.MessageHelper, ctx *models.MessageContext, result *CheckResult, parameter any) bool {

	if MePtn.MatchString(helper.FullName()) {
		result.MarkDelete = true
		result.Message = config.GetTextLang().ErrNameBlock
		return false
	}
	return true
}

package checker

import (
	"encoding/json"
	"gotgpeon/config"
	"gotgpeon/data/models"
	"regexp"
)

var MePtn = regexp.MustCompile("[\u0600-\u06FF\u0750-\u077F\uFB50-\uFDFF]")

func (c *MessageChecker) CheckContentNoMelang(ctx *models.MessageContext, result *CheckResult, parameter json.RawMessage) bool {
	if ctx.Message.Text != "" {
		// check middle east language.
		if MePtn.MatchString(ctx.Message.Text) {
			result.MarkDelete = true
			result.Message = config.GetTextLang().ErrContentNozhtw
			return false
		}
	}
	return true
}

func (c *MessageChecker) CheckNameNoMelang(ctx *models.MessageContext, result *CheckResult, parameter json.RawMessage) bool {

	if MePtn.MatchString(ctx.Message.FullName()) {
		result.MarkDelete = true
		result.Message = config.GetTextLang().ErrNameBlock
		return false
	}
	return true
}

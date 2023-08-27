package checker

import (
	"gotgpeon/models"
	"gotgpeon/pkg/tgbot/core"
)

type CheckerHandler interface {
	CheckMessage(helper *core.MessageHelper, ctx *models.MessageContext) *CheckResult
}

type CheckResult struct {
	MarkDelete bool
	MarkRecord bool
	Message    string
}

type MessageChecker struct {
	checkerMap map[string]func(
		helper *core.MessageHelper,
		ctx *models.MessageContext,
		result *CheckResult,
		parameter interface{},
	) bool
}

func (c *MessageChecker) Init() {
	c.checkerMap = map[string]func(helper *core.MessageHelper, ctx *models.MessageContext, result *CheckResult, parameter interface{}) bool{
		"Type":          c.CheckTypeNoMedia,
		"Forward":       c.CheckNoForward,
		"Entities":      c.CheckEntitiesOK,
		"Viabot":        c.CheckViabotOK,
		"SpchName":      c.CheckNameNospchLang,
		"SpchContent":   c.CheckContentNoSpchLang,
		"MeLangName":    c.CheckNameNoMelang,
		"MeLangContent": c.CheckContentNoMelang,
	}
}

func (c *MessageChecker) CheckMessage(helper *core.MessageHelper, ctx *models.MessageContext) *CheckResult {
	result := &CheckResult{
		MarkDelete: false,
		MarkRecord: true,
		Message:    "",
	}

	if ctx.IsAdminstrator() || ctx.Record.MemberLevel >= models.JUNIOR {
		return result
	}

	// Check message need to delete?
	for _, cfg := range ctx.ChatCfg.CheckerList {
		if checkFunc, ok := c.checkerMap[cfg.Name]; ok {
			isNext := checkFunc(helper, ctx, result, cfg.Parameter)
			if !isNext {
				break
			}
		}
	}

	// Check message need to record?
	if result.MarkDelete {
		result.MarkRecord = false
	}

	return result
}

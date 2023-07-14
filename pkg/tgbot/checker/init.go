package checker

import (
	"gotgpeon/models"
	"gotgpeon/utils"
)

type CheckerHandler interface {
	CheckMessage(helper *utils.MessageHelper, ctx *models.MessageContext) *CheckResult
}

type CheckResult struct {
	MarkDelete bool
	MarkRecord bool
	Message    string
}

type MessageChecker struct {
	checkerMap map[string]func(
		helper *utils.MessageHelper,
		ctx *models.MessageContext,
		result *CheckResult,
		parameter interface{},
	) bool
}

func (c *MessageChecker) Init() {
	c.checkerMap = map[string]func(helper *utils.MessageHelper, ctx *models.MessageContext, result *CheckResult, parameter interface{}) bool{
		"type":     c.CheckTypeOK,
		"entities": c.CheckEntitiesOK,
		"viabot":   c.CheckViabotOK,
		"username": c.CheckNameOK,
		"content":  c.CheckContentOK,
	}
}

func (c *MessageChecker) CheckMessage(helper *utils.MessageHelper, ctx *models.MessageContext) *CheckResult {
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

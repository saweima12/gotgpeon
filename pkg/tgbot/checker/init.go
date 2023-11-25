package checker

import (
	"encoding/json"
	"gotgpeon/data/models"
)

type CheckerHandler interface {
	CheckMessage(ctx *models.MessageContext) *CheckResult
}

type CheckResult struct {
	MarkDelete bool
	MarkRecord bool
	Message    string
}

type MessageChecker struct {
	checkerMap map[string]CheckHandlerFunc
}

type CheckHandlerFunc func(ctx *models.MessageContext, result *CheckResult, parameter json.RawMessage) bool

func (c *MessageChecker) Init() {
	c.checkerMap = map[string]CheckHandlerFunc{
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

func (c *MessageChecker) CheckMessage(ctx *models.MessageContext) *CheckResult {
	result := &CheckResult{
		MarkDelete: false,
		MarkRecord: true,
		Message:    "",
	}

	if ctx.IsAdminstrator() || ctx.Record.MemberLevel >= models.JUNIOR {
		return result
	}

	// Check message need to delete?
	for _, checkFunc := range c.checkerMap {
		isNext := checkFunc(ctx, result, []byte{})
		if !isNext {
			break
		}
	}

	// Check message need to record?
	if result.MarkDelete {
		result.MarkRecord = false
	}

	return result
}

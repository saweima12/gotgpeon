package checker

import (
	"gotgpeon/models"
	"gotgpeon/utils"
)

type CheckerHandler interface {
	CheckMessage(helper *utils.MessageHelper, ctx *models.MessageContext) *CheckResult
}

type CheckResult struct {
	MustDelete bool
	MustRecord bool
	Message    string
}

type MessageChecker struct {
	checkerMap map[string]func(helper *utils.MessageHelper, ctx *models.MessageContext, parameter interface{}) bool
}

func (c *MessageChecker) Init() {
	c.checkerMap = map[string]func(helper *utils.MessageHelper, ctx *models.MessageContext, parameter interface{}) bool{
		"type":     c.CheckTypeOK,
		"entities": c.CheckEntitiesOK,
		"viabot":   c.CheckViabotOK,
	}
}

func (c *MessageChecker) CheckMessage(helper *utils.MessageHelper, ctx *models.MessageContext) *CheckResult {

	result := CheckResult{
		MustDelete: false,
		MustRecord: true,
		Message:    "",
	}

	// Check message need to delete?
	for _, cfg := range ctx.ChatCfg.CheckerList {
		if checkFunc, ok := c.checkerMap[cfg.Name]; ok {
			checkOK := checkFunc(helper, ctx, cfg.Parameter)
			if !checkOK {
				result.MustDelete = true
				break
			}
		}
	}

	// Check message need to record?
	if result.MustDelete {
		result.MustRecord = false
	}

	return &result
}

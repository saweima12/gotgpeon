package checker

import (
	"fmt"
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
		"type": c.CheckType,
	}
}

func (c *MessageChecker) Test(helper *utils.MessageHelper, ctx *models.MessageContext, parameter any) bool {
	return false
}

func (c *MessageChecker) CheckMessage(helper *utils.MessageHelper, ctx *models.MessageContext) *CheckResult {

	result := CheckResult{
		MustDelete: false,
		MustRecord: true,
		Message:    "",
	}

	fmt.Println(ctx.ChatCfg)

	// // Check message need to delete?
	for _, cfg := range ctx.ChatCfg.CheckerList {
		fmt.Println(cfg)
		checkFunc, ok := c.checkerMap[cfg.Name]
		if ok {
			fmt.Println("ok")
			mustDelete := checkFunc(helper, ctx, cfg.Parameter)
			if mustDelete {
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

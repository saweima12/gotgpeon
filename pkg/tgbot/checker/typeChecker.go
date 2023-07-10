package checker

import (
	"fmt"
	"gotgpeon/logger"
	"gotgpeon/models"
	"gotgpeon/utils"
)

type CheckTypeParams struct {
	Data string `json:"data"`
}

func (c *MessageChecker) CheckType(helper *utils.MessageHelper, ctx *models.MessageContext, parameter any) bool {
	var params CheckTypeParams
	err := utils.AnyToStruct(parameter, &params)
	if err != nil {
		logger.Error(err)
	}

	fmt.Println(params)

	return false

}

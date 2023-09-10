package checker

import (
	"gotgpeon/config"
	"gotgpeon/libs/gocc"
	"gotgpeon/data/models"
	"gotgpeon/pkg/tgbot/core"
	"regexp"
	"strings"
)

var ChPtn = regexp.MustCompile("[\u3400-\u4DBF\u4E00-\u9FFF\uF900-\uFAFF]")

func (c *MessageChecker) CheckContentNoSpchLang(helper *core.MessageHelper, ctx *models.MessageContext, result *CheckResult, parameter any) bool {
	if helper.Text != "" {
		// check spchinese
		if !checkSpChineseOK(helper.Text, 2) {
			result.MarkDelete = true
			result.Message = config.GetTextLang().ErrContentNozhtw
			return false
		}
	}

	return true
}

func (c *MessageChecker) CheckNameNospchLang(helper *core.MessageHelper, ctx *models.MessageContext, result *CheckResult, parameter any) bool {
	if !checkSpChineseOK(helper.FullName(), 1) {
		// Check sender name
		result.MarkDelete = true
		result.Message = config.GetTextLang().ErrNameBlock
		return false
	}

	return true
}

func checkSpChineseOK(text string, limit int) bool {
	point := 0
	// Find all chinese char by regex.
	allCh := ChPtn.FindAllString(text, -1)
	// Merge to a string
	originStr := strings.Join(allCh, "")
	tcStr := gocc.S2T(originStr)

	originStrRunes := []rune(originStr)
	tcStrRunes := []rune(tcStr)

	ignoreWordMap := config.GetIgnoreWordMap()
	// Compare string
	for index, chStrRune := range originStrRunes {
		// is ignore word ? skipt it.
		if _, ok := ignoreWordMap[string(chStrRune)]; ok {
			continue
		}

		if originStrRunes[index] != tcStrRunes[index] {
			point += 1
		}

		if point >= limit {
			return false
		}
	}
	return true
}

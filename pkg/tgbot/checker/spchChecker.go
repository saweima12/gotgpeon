package checker

import (
	"gotgpeon/config"
	"gotgpeon/models"
	"gotgpeon/utils"
	"gotgpeon/utils/goccutil"
	"regexp"
	"strings"
)

var MePtn = regexp.MustCompile("[\u0600-\u06FF\u0750-\u077F\uFB50-\uFDFF]")
var ChPtn = regexp.MustCompile("[\u3400-\u4DBF\u4E00-\u9FFF\uF900-\uFAFF]")

func (c *MessageChecker) CheckContentOK(helper *utils.MessageHelper, ctx *models.MessageContext, result *CheckResult, parameter any) bool {
	if helper.Text != "" {
		// check middle east language.
		if MePtn.MatchString(helper.Text) {
			result.MarkDelete = true
			result.Message = config.GetTextLang().TipsContentNozhtw
			return false
		}

		// check spchinese
		if !checkSpChineseOK(helper.Text, 2) {
			result.MarkDelete = true
			result.Message = config.GetTextLang().TipsContentNozhtw
			return false
		}
	}

	return true
}

func (c *MessageChecker) CheckNameOK(helper *utils.MessageHelper, ctx *models.MessageContext, result *CheckResult, parameter any) bool {

	// check middle east language.
	if MePtn.MatchString(helper.FullName()) {
		result.MarkDelete = true
		result.Message = config.GetTextLang().TipsNameBlock
		return false
	}

	if !checkSpChineseOK(helper.FullName(), 1) {
		result.MarkDelete = true
		result.Message = config.GetTextLang().TipsNameBlock
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
	tcStr := goccutil.S2T(originStr)

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

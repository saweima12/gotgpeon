package models

import "gotgpeon/utils/sliceutil"

type MessageContext struct {
	ChatCfg     *ChatConfig `json:"chat_cfg"`
	IsWhitelist bool        `json:"is_whitelist"`
	Point       int         `json:"point"`
}

func (ctx *MessageContext) IsAdminstrator(userId string) bool {
	return sliceutil.Contains(userId, ctx.ChatCfg.AdminStrators)
}

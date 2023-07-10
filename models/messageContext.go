package models

import "gotgpeon/utils/sliceutil"

type MessageContext struct {
	ChatCfg     *ChatConfig    `json:"chat_cfg"`
	IsWhitelist bool           `json:"is_whitelist"`
	Record      *MessageRecord `json:"record"`
}

func (ctx *MessageContext) IsAdminstrator(userId string) bool {
	return sliceutil.Contains(userId, ctx.ChatCfg.Adminstrators)
}

package models

import (
	"gotgpeon/config"
	"gotgpeon/utils/sliceutil"
)

type MessageContext struct {
	ChatCfg     *ChatConfig          `json:"chat_cfg"`
	CommonCfg   *config.CommonConfig `json:"common_cfg"`
	IsWhitelist bool                 `json:"is_whitelist"`
	Record      *MessageRecord       `json:"record"`
}

func (ctx *MessageContext) IsAdminstrator() bool {
	return sliceutil.Contains(ctx.Record.MemberId, ctx.ChatCfg.Adminstrators)
}

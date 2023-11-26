package models

import (
	"gotgpeon/config"
	"gotgpeon/pkg/tgbot/core"
	"gotgpeon/utils/sliceutil"
)

type MessageContext struct {
	ChatCfg     *ChatConfig          `json:"chat_cfg"`
	CommonCfg   *config.CommonConfig `json:"common_cfg"`
	Message     *core.MessageHelper  `json:"message"`
	IsAllowlist bool                 `json:"is_allowlist`
	Record      *MessageRecord       `json:"record"`
}

func (ctx *MessageContext) IsAdminstrator() bool {
	return sliceutil.Contains(ctx.Record.MemberId, ctx.ChatCfg.Adminstrators)
}

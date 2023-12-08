package models

import (
	"encoding/json"
	"gotgpeon/utils/sliceutil"
)

type CheckerConfig struct {
	Name      string          `json:"name"`
	Parameter json.RawMessage `json:"parameter,omitempty"`
}

type ChatConfig struct {
	Status        int             `json:"status"`
	ChatId        int64           `json:"chat_id"`
	ChatName      string          `json:"chat_name"`
	Adminstrators []int64         `json:"adminstrators"`
	CheckerList   []CheckerConfig `json:"checker_config"`
}

func NewDefaultChatConfig(chatId int64, adminstrators []int64) *ChatConfig {
	return &ChatConfig{
		Status:        NG,
		ChatId:        chatId,
		Adminstrators: adminstrators,
		CheckerList:   []CheckerConfig{},
	}
}

func (c *ChatConfig) IsAvaliable() bool {
	return c.Status == OK
}

func (c *ChatConfig) IsAdminstrator(memberId int64) bool {
	return sliceutil.Contains(memberId, c.Adminstrators)
}

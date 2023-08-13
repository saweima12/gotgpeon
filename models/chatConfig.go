package models

import "gotgpeon/utils/sliceutil"

type CheckerConfig struct {
	Name      string      `json:"name"`
	Parameter interface{} `json:"parameter,omitempty"`
}

type ChatConfig struct {
	Status        int             `json:"status"`
	ChatId        int64           `json:"chat_id"`
	ChatName      string          `json:"chat_name"`
	Adminstrators []int64         `json:"adminstrators"`
	CheckerList   []CheckerConfig `json:"checker_config"`
}

func NewDefaultChatConfig(chatId int64, chatName string, adminstrators []int64) *ChatConfig {
	return &ChatConfig{
		Status:        NG,
		ChatId:        chatId,
		ChatName:      chatName,
		Adminstrators: adminstrators,
		CheckerList: []CheckerConfig{
			{Name: "Type"},
			{Name: "Forward"},
			{Name: "Entities"},
			{Name: "Viabot"},
			{Name: "SpchName"},
			{Name: "SpchContent"},
		},
	}
}

func (c *ChatConfig) IsAvaliable() bool {
	return c.Status == OK
}

func (c *ChatConfig) IsAdminstrator(memberId int64) bool {
	return sliceutil.Contains(memberId, c.Adminstrators)
}

package models

import "gotgpeon/utils/sliceutil"

type CheckerConfig struct {
	Name      string      `json:"name"`
	Parameter interface{} `json:"parameter,omitempty"`
}

type ChatConfig struct {
	Status           int             `json:"status"`
	ChatId           int64           `json:"chat_id"`
	ChatName         string          `json:"chat_name"`
	SeniorCount      int             `json:"senior_count"`
	JuniorDay        int             `json:"junior_day"`
	SeniorDay        int             `json:"senior_day"`
	CheckLowestCount int             `json:"check_lowest_count"`
	Adminstrators    []int64         `json:"adminstrators"`
	CheckerList      []CheckerConfig `json:"checker_config"`
}

func NewDefaultChatConfig(chatId int64, chatName string, adminstrators []int64) *ChatConfig {
	return &ChatConfig{
		Status:           NG,
		ChatId:           chatId,
		ChatName:         chatName,
		SeniorCount:      300,
		JuniorDay:        7,
		SeniorDay:        60,
		CheckLowestCount: 20,
		Adminstrators:    adminstrators,
		CheckerList: []CheckerConfig{
			{Name: "type"},
			{Name: "entities"},
			{Name: "viabot"},
			{Name: "username"},
			{Name: "content"},
		},
	}
}

func (c *ChatConfig) IsAvaliable() bool {
	return c.Status == OK
}

func (c *ChatConfig) IsAdminstrator(memberId int64) bool {
	return sliceutil.Contains(memberId, c.Adminstrators)
}

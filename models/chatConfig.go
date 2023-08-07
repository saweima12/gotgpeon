package models

import "gotgpeon/utils/sliceutil"

type CheckerConfig struct {
	Name      string      `json:"name"`
	Parameter interface{} `json:"parameter,omitempty"`
}

type ChatConfig struct {
	Status           string          `json:"status"`
	ChatId           string          `json:"chat_id"`
	ChatName         string          `json:"chat_name"`
	SeniorCount      int             `json:"senior_count"`
	JuniorDay        int             `json:"junior_day"`
	SeniorDay        int             `json:"senior_day"`
	CheckLowestCount int             `json:"check_lowest_count"`
	Adminstrators    []string        `json:"adminstrators"`
	CheckerList      []CheckerConfig `json:"checker_config"`
}

func NewDefaultChatConfig(chatId string, chatName string, adminstrators []string) *ChatConfig {
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

func (c *ChatConfig) IsAdminstrator(idStr string) bool {
	return sliceutil.Contains(idStr, c.Adminstrators)
}

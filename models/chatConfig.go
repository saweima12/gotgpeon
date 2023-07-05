package models

type ChatConfig struct {
	Status           string   `json:"status"`
	ChatId           string   `json:"chat_id"`
	ChatName         string   `json:"chat_name"`
	SeniorCount      int      `json:"senior_count"`
	JuniorDay        int      `json:"junior_day"`
	SeniorDay        int      `json:"senior_day"`
	CheckLowestCount int      `json:"check_lowest_count"`
	AdminStrators    []string `json:"admin_strators"`
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
		AdminStrators:    adminstrators,
	}
}

func (c *ChatConfig) IsAvaliable() bool {
	return c.Status == OK
}

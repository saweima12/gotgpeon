package models

type ChatConfig struct {
	Status           string
	ChatName         string
	SeniorCount      int
	JuniorDay        int
	SeniorDay        int
	CheckLowestCount int
	AdminStrators    []string
}

func NewDefaultChatConfig(chatName string, adminstrators []string) *ChatConfig {
	return &ChatConfig{
		Status:           OK,
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

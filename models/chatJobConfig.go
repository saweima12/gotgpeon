package models

type ChatJobConfig struct {
	SeniorCount      uint
	JuniorDay        uint
	SeniorDay        uint
	CheckLowestCount uint
}

func NewDefaultChatJobConfig() *ChatJobConfig {
	return &ChatJobConfig{
		SeniorCount:      300,
		JuniorDay:        7,
		SeniorDay:        90,
		CheckLowestCount: 20,
	}
}

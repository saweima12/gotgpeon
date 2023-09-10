package models

type ChatJobConfig struct {
	JuniorLowest     uint
	JuniorDay        uint
	SeniorDay        uint
	CheckLowestCount uint
}

func NewDefaultChatJobConfig() *ChatJobConfig {
	return &ChatJobConfig{
		JuniorLowest:     300,
		JuniorDay:        7,
		SeniorDay:        90,
		CheckLowestCount: 20,
	}
}

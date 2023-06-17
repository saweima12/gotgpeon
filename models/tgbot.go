package models

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TgBotClient struct {
	BotAPI     *tgbotapi.BotAPI
	QuitChan   chan bool
	UpdateChan tgbotapi.UpdatesChannel
}

func (c *TgBotClient) Stop() {
	c.QuitChan <- true
}

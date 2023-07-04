package models

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TgbotUpdateProcess struct {
	BotAPI     *tgbotapi.BotAPI
	QuitChan   chan bool
	UpdateChan tgbotapi.UpdatesChannel
}

func (c *TgbotUpdateProcess) Stop() {
	if c.QuitChan != nil {
		c.QuitChan <- true
	}

	if c.BotAPI != nil {
		c.BotAPI.StopReceivingUpdates()
	}
}

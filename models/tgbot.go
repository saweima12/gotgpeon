package models

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TgbotUpdateProcess struct {
	BotAPI     *tgbotapi.BotAPI
	QuitChan   chan struct{}
	UpdateChan tgbotapi.UpdatesChannel
}

func (c *TgbotUpdateProcess) Stop() {
	if c.QuitChan != nil {
		c.QuitChan <- struct{}{}
	}

	if c.BotAPI != nil {
		c.BotAPI.StopReceivingUpdates()
	}
}

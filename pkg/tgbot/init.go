package tgbot

import (
	"fmt"
	"gotgpeon/config"
	"gotgpeon/models"
	"gotgpeon/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func InitTgBot(cfg *config.TgBotConfig) (*tgbotapi.BotAPI, error) {
	// Create telegram bot instance.
	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		return nil, err
	}

	if cfg.AutoSetWebhook {
		err = SetWebhook(cfg, bot)
		if err != nil {
			return nil, err
		}
	}

	// Get update channel.
	ch := bot.ListenForWebhook("/" + bot.Token)

	// Tgbot Client
	client := models.TgBotClient{
		BotAPI:     bot,
		QuitChan:   make(chan bool),
		UpdateChan: ch,
	}

	go ProcessUpdate(&client)
	return bot, nil
}

func ProcessUpdate(c *models.TgBotClient) {
	for {
		select {
		case <-c.QuitChan:
			return
		case update := <-c.UpdateChan:
			fmt.Println(update)
		default:
		}
	}
}

func SetWebhook(cfg *config.TgBotConfig, bot *tgbotapi.BotAPI) error {
	// Try set webhook
	webhook, err := tgbotapi.NewWebhook(cfg.HookURL)
	if err != nil {
		return err
	}

	// Send request to telegram api
	resp, err := bot.Request(webhook)
	if err != nil {
		return err
	}

	if !resp.Ok {
		return utils.ErrSetWebHook
	}

	return nil
}

package tgbot

import (
	"fmt"
	"gotgpeon/config"
	"gotgpeon/pkg/tgbot/handler"

	"gopkg.in/telebot.v3"
)

func InitTgBot(cfg *config.TgBotConfig) (*telebot.Bot, error) {

	// define webhook
	webhook := telebot.Webhook{
		Listen: "3000",
	}

	engine, err := telebot.NewBot(telebot.Settings{
		Token:   cfg.BotToken,
		Offline: true,
		Poller:  &webhook,
		Verbose: true,
	})
	if err != nil {
		return nil, err
	}

	// Initialize Handler
	handler.InitHandler(engine)

	fmt.Println(engine)
	return engine, nil
}

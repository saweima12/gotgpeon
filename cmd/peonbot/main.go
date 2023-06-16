package main

import (
	"flag"
	"fmt"
	"gotgpeon/config"
	"gotgpeon/db"
	"gotgpeon/logger"
	"gotgpeon/pkg/tgbot"
)

var configPath string

func main() {
	// Get configPath parameter.
	inputConfigPath := flag.String("configPath", "config.yml", "configuration file path.")
	flag.Parse()

	// loading configuration
	configPath = *inputConfigPath
	err := config.InitConfig(configPath)
	if err != nil {
		panic("Loading config error path:" + configPath)
	}

	conf := config.GetConfig()
	// Initialize Logger
	err = logger.InitLogger()

	if err != nil {
		panic("Initialize Logger err:" + err.Error())
	}

	// Initialize database and redis connection.
	err = db.InitDbConn(&conf.Common)
	if err != nil {
		panic("Initialize Database connection err:" + err.Error())
	}

	// Initialize Bot
	bot, err := tgbot.InitTgBot(&config.GetConfig().TgBot)
	if err != nil {
		panic("Initialize telegram bot err:" + err.Error())
	}

	logger.Info("Initialize finished.")
	fmt.Println(bot.Poller)
	bot.Start()
}

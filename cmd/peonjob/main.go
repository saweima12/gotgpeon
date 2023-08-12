package main

import (
	"flag"
	"gotgpeon/config"
	"gotgpeon/db"
	"gotgpeon/logger"
	"gotgpeon/pkg/schedule"
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

	cfg := config.GetConfig()
	// Initialize Logger
	err = logger.InitLogger(cfg.Common.Mode)
	if err != nil {
		panic("Initialize Logger err:" + err.Error())
	}

	// Initialize database and redis connection.
	err = db.InitDbConn(&cfg.Common)
	if err != nil {
		panic("Initialize Database connection err:" + err.Error())
	}

	// Initialize BotAPI
	botClient, err := tgbot.InitTgBot(&config.GetConfig().TgBot)
	if err != nil {
		panic("Initialize telegram bot err:" + err.Error())
	}

	// Initialize Schedule.
	peonSchedule, err := schedule.NewPeonSchedule(botClient)
	if err != nil {
		panic("Initialize schedule err:" + err.Error())
	}
	peonSchedule.Run()

}

package main

import (
	"flag"
	"gotgpeon/config"
	"gotgpeon/db"
	"gotgpeon/logger"
	"gotgpeon/pkg/tgbot"
	"net/http"
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
	err = logger.InitLogger()
	if err != nil {
		panic("Initialize Logger err:" + err.Error())
	}

	// Initialize database and redis connection.
	err = db.InitDbConn(&cfg.Common)
	if err != nil {
		panic("Initialize Database connection err:" + err.Error())
	}

	// Initialize Bot & Add webhook
	botClient, err := tgbot.InitTgBot(&config.GetConfig().TgBot)
	if err != nil {
		panic("Initialize telegram bot err:" + err.Error())
	}

	// Add Webhook route and launche update process.
	client := tgbot.StartUpdateProcess(cfg.TgBot.BotToken, botClient)
	// When shutdown timing, close the UpdateProcess
	defer client.Stop()
	logger.Info("Initialize finished.")

	// Start a http server for listen update
	err = http.ListenAndServe(cfg.Common.ListenPort, nil)
	if err != nil {
		logger.Error("Listen hook err" + err.Error())
	}

}

package main

import (
	"flag"
	"fmt"
	"gotgpeon/config"
	"gotgpeon/db"
	"gotgpeon/logger"
	"gotgpeon/models"
	"gotgpeon/pkg/tgbot"
	"gotgpeon/utils/goccutil"
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
	err = logger.InitLogger(cfg.Common.Mode)
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

	var client *models.TgbotUpdateProcess
	// Add Webhook route and launche update process.
	if cfg.Common.Mode == "webhook" {
		client = tgbot.StartWebhookProcess(cfg.TgBot.BotToken, botClient)
		fmt.Println("UpdateMode: Webhook")
	} else {
		client = tgbot.StartLongPollProcess(botClient)
		fmt.Println("UpdateMode: LongPoll")
	}

	// When shutdown timing, close the UpdateProcess
	defer client.Stop()
	// Initialize opencc
	goccutil.InitOpenCC()

	logger.Info("Initialize finished.")

	// Start a http server for listen update
	err = http.ListenAndServe(":"+cfg.Common.ListenPort, nil)
	if err != nil {
		logger.Error("Listen hook err" + err.Error())
	}

}

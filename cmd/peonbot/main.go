package main

import (
	"flag"
	"gotgpeon/config"
	"gotgpeon/db"
	"gotgpeon/libs/ants"
	"gotgpeon/libs/gocc"
	"gotgpeon/libs/timewheel"
	"gotgpeon/logger"
	"gotgpeon/models"
	"gotgpeon/pkg/tgbot"
	"gotgpeon/utils"
	"net/http"
)

var configPath string
var langPath string

func main() {
	// Get configPath parameter.
	inputConfigPath := flag.String("configPath", "config.yml", "configuration file path.")
	inputLangPath := flag.String("langPath", "lang.yml", "language file path.")
	flag.Parse()

	// loading configuration
	configPath = *inputConfigPath
	err := config.InitConfig(configPath)
	if err != nil {
		panic("Loading config error path:" + configPath)
	}

	// loading language file.
	langPath = *inputLangPath
	err = config.InitTextlang(langPath)
	if err != nil {
		panic("Loading language file error, path: " + langPath + " err:" + err.Error())
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

	// Initialize ants pool.
	err = ants.Init()
	if err != nil {
		panic("Initialize goroutine pool err" + err.Error())
	}

	// Initialize timingwheel.
	wheel, err := timewheel.Init()
	if err != nil {
		panic("Initialize timewheel err: %s" + err.Error())
	}
	defer wheel.Stop()

	// Initialize Bot & Add webhook
	botClient, err := tgbot.InitTgBot(&config.GetConfig().TgBot)
	if err != nil {
		panic("Initialize telegram bot err:" + err.Error())
	}
	if cfg.Common.Mode == "dev" {
		tgbot.AttachTestMe(botClient)
	}
	defer tgbot.DeleteWebhook(botClient)

	var client *models.TgbotUpdateProcess
	// Add Webhook route and launche update process.
	if cfg.TgBot.UpdateMode == "webhook" {
		client = tgbot.StartWebhookProcess(cfg.TgBot.BotToken, botClient)
		logger.Info("UpdateMode: Webhook")
	} else {
		tgbot.DeleteWebhook(botClient)
		client = tgbot.StartLongPollProcess(botClient)
		logger.Info("UpdateMode: LongPoll")
	}
	// When shutdown timing, close the UpdateProcess
	defer client.Stop()

	// Initialize opencc
	gocc.InitOpenCC()
	logger.Info("Initialize finished.")

	// handle signal
	c := utils.HandleSingal()

	// Start a http server for listen update
	go func() {
		err = http.ListenAndServe(":"+cfg.Common.ListenPort, http.DefaultServeMux)
		if err != nil {
			logger.Error("Listen hook err" + err.Error())
		}
	}()

	<-c
	logger.Info("Service shutdown...")
}

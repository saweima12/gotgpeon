package main

import (
	"flag"
	"gotgpeon/config"
	"gotgpeon/db"
	"gotgpeon/logger"
	"gotgpeon/models"
	"gotgpeon/pkg/tgbot"
	"gotgpeon/utils"
	"gotgpeon/utils/goccutil"
	"gotgpeon/utils/poolutil"
	"gotgpeon/utils/timewheel"
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
	err = poolutil.Init()
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

	var client *models.TgbotUpdateProcess
	// Add Webhook route and launche update process.
	if cfg.Common.Mode == "webhook" {
		client = tgbot.StartWebhookProcess(cfg.TgBot.BotToken, botClient)
		logger.Info("UpdateMode: Webhook")
	} else {
		client = tgbot.StartLongPollProcess(botClient)
		logger.Info("UpdateMode: LongPoll")
	}
	// When shutdown timing, close the UpdateProcess
	defer client.Stop()

	// Initialize opencc
	goccutil.InitOpenCC()
	logger.Info("Initialize finished.")

	// handle signal
	c := utils.HandleSingal()

	// Start a http server for listen update
	go func() {
		err = http.ListenAndServe(":"+cfg.Common.ListenPort, nil)
		if err != nil {
			logger.Error("Listen hook err" + err.Error())
		}
	}()

	<-c
	logger.Info("Service shutdown...")
}

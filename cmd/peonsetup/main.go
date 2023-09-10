package main

import (
	"flag"
	"fmt"
	"gotgpeon/config"
	"gotgpeon/db"
	"gotgpeon/logger"
	"gotgpeon/data/entity"
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

	// Initialize Logger
	cfg := config.GetConfig()
	err = logger.InitLogger(cfg.Common.Mode)
	if err != nil {
		panic("Initialize Logger err:" + err.Error())
	}

	// Initialize database and redis connection.
	err = db.InitDbConn(&cfg.Common)
	if err != nil {
		panic("Initialize Database connection err:" + err.Error())
	}

	InitSchema()

	// Initialize webhook
	if cfg.TgBot.UpdateMode != "webhook" {
		return
	}

	botAPI, err := tgbot.InitTgBot(&cfg.TgBot)
	if err != nil {
		panic("Intitialize Telegram bot err: " + err.Error())
	}
	endpoint := cfg.TgBot.HookURL + cfg.TgBot.BotToken
	fmt.Println("Setup Webhook endpoint:" + endpoint)
	err = tgbot.SetWebhook(endpoint, botAPI)
	if err != nil {
		panic("SetWebhook err:" + err.Error())
	}

}

func InitSchema() {
	conn := db.GetDB()
	conn.AutoMigrate(entity.PeonChatConfig{})
	// Create PeonBehaviorRecord and constrains
	conn.AutoMigrate(entity.PeonMemberRecord{})
	conn.AutoMigrate(entity.PeonChatMemberRecord{})
	// conn.Exec("ALTER TABLE public.peon_behavior_record ADD CONSTRAINT peon_behavior_record_un UNIQUE (chat_id,user_id);")
	conn.AutoMigrate(entity.PeonSavedMessage{})
	conn.AutoMigrate(entity.PeonMemberAllowlist{})
	conn.AutoMigrate(entity.PeonDeletedMessage{})
}

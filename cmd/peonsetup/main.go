package main

import (
	"flag"
	"gotgpeon/config"
	"gotgpeon/db"
	"gotgpeon/logger"
	"gotgpeon/models/entity"
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

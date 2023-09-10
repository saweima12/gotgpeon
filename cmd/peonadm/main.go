package main

import (
	"flag"
	"gotgpeon/config"
	"gotgpeon/db"
	"gotgpeon/logger"
	"gotgpeon/pkg/admsite"
	"gotgpeon/utils"

	"github.com/gin-gonic/gin"
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

	// Initialize server & router.
	engine := gin.Default()
	admsite.InitRouter(engine)

	// Handle signal
	c := utils.HandleSingal()

	go engine.Run(":9001")
	<-c

	logger.Info("peonadm shutting down...")
}

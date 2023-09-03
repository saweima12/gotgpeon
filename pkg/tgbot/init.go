package tgbot

import (
	"encoding/json"
	"fmt"
	"gotgpeon/config"
	"gotgpeon/db"
	"gotgpeon/logger"
	"gotgpeon/models"
	"gotgpeon/pkg/tgbot/handler"
	"gotgpeon/utils"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func InitTgBot(cfg *config.TgBotConfig) (*tgbotapi.BotAPI, error) {
	// Create telegram bot instance.
	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		return nil, err
	}

	if cfg.AutoSetWebhook {
		err = SetWebhook(cfg.HookURL, bot)
		if err != nil {
			return nil, err
		}
	}

	return bot, err
}

// Start to revice and handle message from webhook
func StartWebhookProcess(botToken string, botAPI *tgbotapi.BotAPI) *models.TgbotUpdateProcess {
	// Get update channel.
	ch := botAPI.ListenForWebhook("/" + botToken)

	// Get updateProcess instance.
	process := models.TgbotUpdateProcess{
		BotAPI:     botAPI,
		QuitChan:   make(chan struct{}),
		UpdateChan: ch,
	}

	go runUpdateProcess(&process, botAPI)
	return &process
}

// Start to revice and handle message from longpoll
func StartLongPollProcess(botAPI *tgbotapi.BotAPI) *models.TgbotUpdateProcess {

	upCfg := tgbotapi.NewUpdate(0)
	ch := botAPI.GetUpdatesChan(upCfg)

	process := models.TgbotUpdateProcess{
		BotAPI:     botAPI,
		UpdateChan: ch,
	}

	go runUpdateProcess(&process, botAPI)
	return &process
}

func SetWebhook(hookURL string, bot *tgbotapi.BotAPI) error {
	// Try set webhook
	webhook, err := tgbotapi.NewWebhook(hookURL)
	if err != nil {
		return err
	}
	// Send request to telegram api
	resp, err := bot.Request(webhook)
	if err != nil {
		return err
	}

	if !resp.Ok {
		return utils.ErrSetWebHook
	}

	return nil
}

func DeleteWebhook(bot *tgbotapi.BotAPI) error {
	cfg := tgbotapi.DeleteWebhookConfig{DropPendingUpdates: false}

	resp, err := bot.Request(cfg)
	if err != nil {
		return err
	}

	fmt.Println(resp)
	return nil
}

func ProcessUpdate(msgHandler handler.MessageHandler, update tgbotapi.Update, botAPI *tgbotapi.BotAPI) {
	if update.Message != nil {
		msgHandler.HandleMessage(update.Message, botAPI, false)
	}
	if update.EditedMessage != nil {
		msgHandler.HandleMessage(update.EditedMessage, botAPI, true)
	}
}

func runUpdateProcess(c *models.TgbotUpdateProcess, botAPI *tgbotapi.BotAPI) {
	dbConn := db.GetDB()
	cacheConn := db.GetCache()
	// Create handler.
	msgHandler := handler.NewMessageHandler(dbConn, cacheConn, botAPI)
LOOP:
	for {
		select {
		case <-c.QuitChan:
			break LOOP
		case update := <-c.UpdateChan:
			ProcessUpdate(msgHandler, update, botAPI)
		}
	}

	logger.Info("UpdateProcess stop ...")
}

func AttachTestMe(botAPI *tgbotapi.BotAPI) {
	// Add test function.
	http.HandleFunc("/"+botAPI.Token+"/me", func(w http.ResponseWriter, r *http.Request) {
		me, _ := botAPI.GetMe()
		hookInfo, _ := botAPI.GetWebhookInfo()

		result := make(map[string]interface{})
		result["Me"] = me
		result["Hook"] = hookInfo

		bytes, _ := json.Marshal(result)
		w.Write(bytes)
	})
}

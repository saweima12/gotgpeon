package tgbot

import (
	"gotgpeon/config"
	"gotgpeon/db"
	"gotgpeon/models"
	"gotgpeon/pkg/tgbot/handler"
	"gotgpeon/utils"
	"gotgpeon/utils/poolutil"

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
		QuitChan:   make(chan bool),
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

func ProcessUpdate(msgHandler handler.MessageHandler, update tgbotapi.Update, botAPI *tgbotapi.BotAPI) func() {
	return func() {
		if update.Message != nil {
			msgHandler.HandleMessage(update.Message, botAPI, false)
		}
		if update.EditedMessage != nil {
			msgHandler.HandleMessage(update.EditedMessage, botAPI, true)
		}
	}
}

func runUpdateProcess(c *models.TgbotUpdateProcess, botAPI *tgbotapi.BotAPI) {
	dbConn := db.GetDB()
	cacheConn := db.GetCache()
	// Create handler.
	msgHandler := handler.NewMessageHandler(dbConn, cacheConn, botAPI)

	for {
		select {
		case <-c.QuitChan:
			return
		case update := <-c.UpdateChan:
			{
				poolutil.Submit(ProcessUpdate(msgHandler, update, botAPI))
			}
		default:
		}
	}
}

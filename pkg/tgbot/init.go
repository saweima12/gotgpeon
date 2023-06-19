package tgbot

import (
	"gotgpeon/config"
	"gotgpeon/db"
	"gotgpeon/models"
	"gotgpeon/pkg/tgbot/handler"
	"gotgpeon/utils"

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

func StartUpdateProcess(botToken string, botAPI *tgbotapi.BotAPI) *models.TgbotUpdateProcess {
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

func ProcessUpdate(msgHandler handler.MessageHandler, update tgbotapi.Update, botAPI *tgbotapi.BotAPI) {

	if update.Message != nil {
		msgHandler.HandleMessage(update.Message, botAPI)
	}

	if update.EditedMessage != nil {
		msgHandler.HandleMessage(update.EditedMessage, botAPI)
	}

}

func runUpdateProcess(c *models.TgbotUpdateProcess, botAPI *tgbotapi.BotAPI) {

	dbConn := db.GetDB()
	cacheConn := db.GetCache()
	// Create handler.
	msgHandler := handler.NewMessageHandler(dbConn, cacheConn)

	for {
		select {
		case <-c.QuitChan:
			return
		case update := <-c.UpdateChan:
			{
				ProcessUpdate(msgHandler, update, botAPI)
			}
		default:
		}
	}
}

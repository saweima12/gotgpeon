package handler

import "gopkg.in/telebot.v3"

func InitHandler(engine *telebot.Bot) {
	engine.Handle(telebot.OnEdited, func(context telebot.Context) error {
		return nil
	})
}

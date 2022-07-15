package main

import (
	tele "gopkg.in/telebot.v3"
)

func (bot Bot) SetUpHandlers() {
	bot.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})
}

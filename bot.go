package main

import (
	"fmt"
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
)

type Bot struct {
	tele.Bot
}

func createBot() (*Bot, error) {
	bot := &Bot{}
	token := os.Getenv("API_KEY")
	poller := bot.MakePoller(os.Getenv("POLLER_MODE"))

	pref := tele.Settings{
		Token:  token,
		Poller: poller,
	}
	telebot, err := tele.NewBot(pref)
	if err != nil {
		return nil, err
	}

	bot.Bot = *telebot
	return bot, nil
}

func (bot Bot) MakePoller(pollerType string) tele.Poller {

	if pollerType == "webhook" {
		port := os.Getenv("PORT")
		route := os.Getenv("ROUTE")
		listen := fmt.Sprintf(":%s/%s", port, route)
		webhook := os.Getenv("WEBHOOK_URL") + route

		return &tele.Webhook{Listen: listen, Endpoint: &tele.WebhookEndpoint{PublicURL: webhook}}
	}

	return &tele.LongPoller{Timeout: 10 * time.Second}
}

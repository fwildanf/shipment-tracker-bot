package main

import (
	"log"

	tele "gopkg.in/telebot.v3"
	env "github.com/joho/godotenv"
)

func init() {
	if err := env.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	bot, err := createBot()

	if err != nil {
		log.Fatal(err)
		return
	}

	bot.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	bot.SetUpHandlers()
	bot.Start()
}

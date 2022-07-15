package main

import (
	"log"

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

	bot.SetUpHandlers()
	bot.Start()
}

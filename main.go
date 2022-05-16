package main

import (
	"log"
	"main.go/src/bot"
)

func main() {
	log.Println("Starting bot")

	// Calls function for starting the connection to Discord bot
	bot.StartBot()
}

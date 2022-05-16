package bot

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type DiscordBot struct {
	Token string
}

func StartBot() {
	log.Println("Initalizing connection")

	b := DiscordBot{
		Token: "DISCORD BOT TOKEN",
	}

	b.connect()
}

func (b *DiscordBot) connect() {
	// Here we start a new Discord session using the provided bot token.
	db, err := discordgo.New("Bot " + b.Token)
	if err != nil {
		log.Println("Error creating Discord session: ", err)
		log.Println("Exiting...")
		time.Sleep(2 * time.Second)
	}

	// register the messageCreate function as a callback for MessageCreate events
	db.AddHandler(messageCreate)
	db.Identify.Intents = discordgo.IntentsGuildMessages

	// Websocket connection to discord
	err = db.Open()
	if err != nil {
		log.Println("Error opening connection: ", err)
		log.Println("Exiting...")
		time.Sleep(2 * time.Second)
	}

	log.Println("Bot started")

	// Runs the bot until ctrl-c is pressed or other system signal
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Safely close the websocket connection
	err = db.Close()
	if err != nil {
		log.Println("Error closing connection: ", err)
		log.Println("Exiting...")
		time.Sleep(2 * time.Second)
	}
}

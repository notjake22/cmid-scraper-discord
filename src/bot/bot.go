package bot

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"main.go/src/scrape"
	"strings"
)

// The function gets called because of the add handler function we set in the connect function
// The bot will be able to read any channel you give it access to
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Just a safeguard to make sure we don't respond to ourselves
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Checking if the message contains the suffix in the message that we want to respond to
	if strings.Contains(m.Content, "!cmid") {
		// just checking if it contains a valid url
		if strings.Contains(m.Content, "https://") {
			log.Println("Valid link detected")

			_, err := s.ChannelMessageSend(m.ChannelID, "Scraping url...")
			if err != nil {
				log.Println("Error sending message", err)
				return
			}
			// splitting the message into an array of strings to parse the url out of it
			url := strings.Split(m.Content, " ")[1]

			// calling the scaper function to scrape the url for the cmid
			cmid, err := scrape.GetCMID(url)
			if err != nil {
				log.Println("Error getting CMID: ", err)
				_, err = s.ChannelMessageSend(m.ChannelID, "Error getting CMID: ")
				if err != nil {
					log.Println("Error sending message: ", err)
					return
				}
				return
			}

			// sending the cmid to the channel
			_, err = s.ChannelMessageSend(m.ChannelID, "CMID: "+cmid)
			if err != nil {
				log.Println("Error sending message: ", err)
				return
			}

		} else {
			log.Println("Invalid link detected")
			_, err := s.ChannelMessageSend(m.ChannelID, "Invalid url")
			if err != nil {
				log.Println("Error sending message: ", err)
				return
			}
		}
	}
}

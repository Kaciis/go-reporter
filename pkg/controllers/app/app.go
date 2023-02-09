package app

import (
	"os"

	"github.com/bwmarrin/discordgo"
)

var App AppType

func Init() {

	token := os.Getenv("TOKEN")
	channelID := os.Getenv("CHANNEL_ID")
	session, err := discordgo.New("Bot " + token)

	if err != nil {
		panic(err)
	}

	meUser, err := session.User("@me")
	if err != nil {
		panic(err)
	}

	botID := meUser.ID

	App = AppType{
		Session:   session,
		ChannelID: channelID,
		BotID:     botID,
	}
}

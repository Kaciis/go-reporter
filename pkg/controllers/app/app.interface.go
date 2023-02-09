package app

import "github.com/bwmarrin/discordgo"

type AppType struct {
	Session   *discordgo.Session
	ChannelID string
	BotID     string
}

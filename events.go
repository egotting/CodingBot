package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == discord.State.User.ID {
		return
	}

	switch {
	case strings.Contains(message.Content, "/hello"):
		discord.ChannelMessageSend(message.ChannelID, "hello world!")
	case strings.Contains(message.Content, "/bye"):
		discord.ChannelMessageSend(message.ChannelID, "bye!")
	}
}

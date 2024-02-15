package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

func checkNilErr(e error) {
	if e != nil {
		log.Fatal("Error message")
	}
}

func StartBot() {

	// create a session
	discord, err := discordgo.New(GetTokenBot())
	checkNilErr(err)

	// add a event handler
	discord.AddHandler(newMessage)

	discord.Open()
	defer discord.Close()

	fmt.Println("bot is running..")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

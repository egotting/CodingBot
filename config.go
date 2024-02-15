package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var BotToken string = "Bot "

func GetTokenBot() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	token := os.Getenv("TOKEN")
	Token := BotToken + token

	return Token
}

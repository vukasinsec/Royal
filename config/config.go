package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config struktura za konfiguraciju
type Config struct {
	BotToken string
	ChatID   string
	APIURL   string
}

// LoadConfig učitava konfiguraciju iz .env fajla
func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Upozorenje: Nema .env fajla, koristi default vrednosti")
	}

	botToken := os.Getenv("BOT_TOKEN")
	chatID := os.Getenv("GROUP_CHAT_ID") // Koristimo ID grupe

	if botToken == "" || chatID == "" {
		log.Fatal("Greška: Nema BOT_TOKEN ili GROUP_CHAT_ID u konfiguraciji")
	}

	return Config{
		BotToken: botToken,
		ChatID:   chatID,
		APIURL:   "https://api.telegram.org/bot" + botToken,
	}
}

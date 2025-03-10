package notify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"rasperon/website-monitor/config"
)

type discordMessage struct {
	Content string `json:"content"`
}

func StartDiscordWorker(discordChan chan string) {
	for msg := range discordChan {
		SendDiscordMessage(msg)
		time.Sleep(2 * time.Second) // Rate limit önlemi
	}
}

func SendDiscordMessage(message string) {
	if config.Webhook == "" {
		fmt.Println("Discord Webhook girilmedi, bildirim atlanıyor.")
		return
	}

	data, _ := json.Marshal(discordMessage{Content: message})
	_, err := http.Post(config.Webhook, "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Discord Webhook gönderilemedi:", err)
	}
}

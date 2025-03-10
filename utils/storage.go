package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
	"rasperon/website-monitor/config"
)

type ConfigData struct {
	URLs     []string `json:"urls"`
	Interval int      `json:"interval"`
	Webhook  string   `json:"webhook"`
}

const configFile = "config.json"

func SaveConfig() {
	data := ConfigData{
		URLs:     config.URLs,
		Interval: int(config.Interval.Seconds()),
		Webhook:  config.Webhook,
	}

	file, err := os.Create(configFile)
	if err != nil {
		fmt.Println("Config kaydedilemedi:", err)
		return
	}
	defer file.Close()

	json.NewEncoder(file).Encode(data)
}

func LoadConfig() bool {
	file, err := os.Open(configFile)
	if err != nil {
		return false
	}
	defer file.Close()

	var data ConfigData
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return false
	}

	config.URLs = data.URLs
	config.Interval = time.Duration(data.Interval) * time.Second
	config.Webhook = data.Webhook

	fmt.Println("Önceki ayarlar yüklendi.")
	return true
}

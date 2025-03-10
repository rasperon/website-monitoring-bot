package config

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

var URLs []string
var Interval time.Duration
var Webhook string

func GetUserInput() {
	// Eğer daha önce kayıtlı config varsa onu yükleyelim
	if loadConfig() {
		return
	}

	reader := bufio.NewReader(os.Stdin)

	// URL'leri al
	fmt.Println("İzlemek istediğiniz web sitelerini girin (boş satır ile bitirin):")
	for {
		fmt.Print("URL: ")
		url, _ := reader.ReadString('\n')
		url = url[:len(url)-1] // \n karakterini temizle
		if url == "" {
			break
		}
		URLs = append(URLs, url)
	}

	// Kontrol süresi
	fmt.Print("Kontrol aralığını (saniye) girin: ")
	intervalStr, _ := reader.ReadString('\n')
	intervalStr = intervalStr[:len(intervalStr)-1]

	interval, err := strconv.Atoi(intervalStr)
	if err != nil || interval <= 0 {
		fmt.Println("Geçersiz giriş, varsayılan olarak 10 saniye ayarlandı.")
		Interval = 10 * time.Second
	} else {
		Interval = time.Duration(interval) * time.Second
	}

	// Discord Webhook
	fmt.Print("Discord Webhook URL'yi girin (zorunlu değil): ")
	Webhook, _ = reader.ReadString('\n')
	Webhook = Webhook[:len(Webhook)-1]

	// Ayarları kaydedelim
	saveConfig()
}

func loadConfig() bool {
	data, err := os.ReadFile("config.json")
	if err != nil {
		return false
	}

	config := struct {
		URLs     []string
		Interval int
		Webhook  string
	}{}

	if err := json.Unmarshal(data, &config); err != nil {
		return false
	}

	URLs = config.URLs
	Interval = time.Duration(config.Interval) * time.Second
	Webhook = config.Webhook
	return true
}

func saveConfig() error {
	config := struct {
		URLs     []string
		Interval int
		Webhook  string
	}{
		URLs:     URLs,
		Interval: int(Interval.Seconds()),
		Webhook:  Webhook,
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile("config.json", data, 0644)
}

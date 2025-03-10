package monitor

import (
	"fmt"
	"net/http"
	"sync"
	"time"
	"rasperon/website-monitor/utils"
)

var statusCache = make(map[string]int) // Site durumu cache
var cacheMutex sync.Mutex              // Haritayı güvenli kullanmak için mutex

func CheckWebsite(url string, wg *sync.WaitGroup, discordChan chan string) {
	defer wg.Done()

	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)

	if err != nil {
		handleStatus(url, 0, discordChan) // 0 = Erişim Hatası
		return
	}
	defer resp.Body.Close()

	handleStatus(url, resp.StatusCode, discordChan)
}

func handleStatus(url string, status int, discordChan chan string) {
	cacheMutex.Lock()
	prevStatus, exists := statusCache[url]
	statusCache[url] = status
	cacheMutex.Unlock()

	msg := fmt.Sprintf("[%s] %s - Status: %d", time.Now().Format(time.RFC3339), url, status)
	fmt.Println(msg)
	utils.LogToFile(msg)

	// Eğer statü değişmişse Discord'a mesaj gönder
	if !exists || prevStatus != status {
		discordChan <- msg
	}
}

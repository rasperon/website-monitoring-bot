package main

import (
	"sync"
	"time"
	"rasperon/website-monitor/config"
	"rasperon/website-monitor/monitor"
	"rasperon/website-monitor/notify"
)

func main() {
	config.GetUserInput()

	if len(config.URLs) == 0 {
		println("İzlenecek site bulunamadı, çıkılıyor...")
		return
	}

	discordChan := make(chan string, 10) // Discord mesajları için kanal
	go notify.StartDiscordWorker(discordChan) // Discord mesaj yöneticisini başlat

	for {
		var wg sync.WaitGroup

		for _, url := range config.URLs {
			wg.Add(1)
			go monitor.CheckWebsite(url, &wg, discordChan)
		}

		wg.Wait()
		time.Sleep(config.Interval)
	}
}

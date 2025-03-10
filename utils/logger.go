package utils

import (
	"fmt"
	"os"
)

func LogToFile(message string) {
	file, err := os.OpenFile("logs/monitor.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Log dosyasına yazılamadı:", err)
		return
	}
	defer file.Close()

	file.WriteString(message + "\n")
}

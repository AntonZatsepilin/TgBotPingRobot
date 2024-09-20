package main

import (
	"fmt"
	"goPingRobot/telegram"
	"goPingRobot/workerpool"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

const (
	INTERVAL        = time.Second * 10
	REQUEST_TIMEOUT = time.Second * 2
	WORKERS_COUNT   = 3
)

var urls = []string{
	"https://gb.com/AntonZatsepilin",
	"https://vk.com/antoshka_zac",
	"https://tlgg.ru/@zzwwmp",
	"https://google.com/",
	"https://golang.org/",
}

func main() {

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN is not set")
	}

	chatIDStr := os.Getenv("TELEGRAM_CHAT_ID")
	if chatIDStr == "" {
		log.Fatal("TELEGRAM_CHAT_ID is not set")
	}

	chatID, err := strconv.ParseInt(chatIDStr, 10, 64)
	if err != nil {
		log.Fatalf("Error converting TELEGRAM_CHAT_ID to int64: %v", err)
	}

	telegram.Init(token, chatID)

	results := make(chan workerpool.Result)
	workerPool := workerpool.New(WORKERS_COUNT, REQUEST_TIMEOUT, results)

	workerPool.Init()

	go generateJobs(workerPool)
	go proccessResults(results)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	workerPool.Stop()
}

func proccessResults(results chan workerpool.Result) {
	go func() {
		for result := range results {
			info := result.Info()
			fmt.Println(info)

			if result.Error != nil {
				log.Printf("Error result: %v", result.Error)
				telegram.SendMessage(info)
			}
		}
	}()
}

func generateJobs(wp *workerpool.Pool) {
	for {
		for _, url := range urls {
			wp.Push(workerpool.Job{URL: url})
		}

		time.Sleep(INTERVAL)
	}

}

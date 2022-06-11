package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	delay, err := strconv.Atoi(getEnv("DELAY", "1"))
	if err != nil {
		log.Fatal("DELAY must be a valid integer.")
	}

	var service Service
	service.Endpoint = getEnv("HTTP_ENDPOINT", "http://127.0.0.1:80")
	service.ServiceKey = os.Getenv("PAGERDUTY_SERVICEKEY")

	for {
		service.runCheck()
		time.Sleep(time.Duration(delay) * time.Second)
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

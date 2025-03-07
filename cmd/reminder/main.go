package main

import (
	"log"

	"github.com/pushinist/pills-taking-reminder-service/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
}

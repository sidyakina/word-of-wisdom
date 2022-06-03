package main

import (
	"github.com/sidyakina/word-of-wisdom/internal/client/tcp"
	"log"
	"math/rand"
	"time"
)

func main() {
	cfg, err := parseConfig()
	if err != nil {
		log.Panicf("failed to parse config: %v", err)
	}

	rand.Seed(time.Now().UnixNano())

	tcp.GetQuotes(cfg.TCPServerAddress, cfg.ReadTimeout, cfg.NumberQuotes)
}

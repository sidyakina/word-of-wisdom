package main

import (
	"github.com/sidyakina/word-of-wisdom/internal/client/tcp"
	"log"
)

func main() {
	cfg, err := parseConfig()
	if err != nil {
		log.Panicf("failed to parse config: %v", err)
	}

	tcp.GetQuotes(cfg.TCPServerAddress, cfg.WaitingMessageTimeout, cfg.NumberQuotes)
}

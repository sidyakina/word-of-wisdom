package main

import (
	"github.com/sidyakina/word-of-wisdom/internal/server/quotes-repo"
	"github.com/sidyakina/word-of-wisdom/internal/server/tcp"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg, err := parseConfig()
	if err != nil {
		log.Panicf("failed to parse config: %v", err)
	}

	rand.Seed(time.Now().UnixNano())
	quotesRepo, err := quotesrepo.New(cfg.QuotesFilePath)
	if err != nil {
		log.Panicf("failed to init quotes repo: %v", err)
	}

	tcpServer := tcp.New(quotesRepo)

	go func() {
		err := tcpServer.Listen(cfg.TCPPort)
		if err != nil {
			log.Panicf("failed to listen tcp on port %v: %v", cfg.TCPPort, err)
		}
	}()

	log.Printf("tcp server started")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	<-ch

	log.Println("got signal, tcp server stopped")
}

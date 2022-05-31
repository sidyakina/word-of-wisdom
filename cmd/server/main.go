package main

import (
	quotesrepo "github.com/sidyakina/word-of-wisdom/internal/server/pkg/quotes-repo"
	"github.com/sidyakina/word-of-wisdom/internal/server/pkg/tcp"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := parseConfig()
	if err != nil {
		log.Panicf("failed to parse config: %v", err)
	}

	quotesRepo := quotesrepo.New()

	tcpServer := tcp.New(quotesRepo)

	go func() {
		err := tcpServer.Listen(cfg.TCPPort)
		if err != nil {
			log.Panicf("failed to listen tcp on port %v: %v", cfg.TCPPort, err)
		}
	}()

	log.Printf("tcp server started")

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGKILL)

	<-ch

	log.Println("got signal, tcp server stopped")
}

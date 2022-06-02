package tcp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/sidyakina/word-of-wisdom/pkg/server/api"
	messagetype "github.com/sidyakina/word-of-wisdom/pkg/server/message-type"
	"log"
	"net"
)

type Server struct {
	quotesRepo QuotesRepo
}

type QuotesRepo interface {
	GetQuote() (string, error)
}

func New(quotesRepo QuotesRepo) *Server {
	return &Server{quotesRepo: quotesRepo}
}

func (s *Server) Listen(port int32) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept new conn: %v", err)

			continue
		}

		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	client := newClient(conn)

	for scanner.Scan() {
		rawMessage := scanner.Bytes()

		message := api.Message{}
		err := json.Unmarshal(rawMessage, &message)
		if err != nil {
			log.Printf("failed to unmarshal request %s: %v", message, err)

			continue
		}

		log.Printf("new message with type %v and payload %s", message.Type, message.Payload)

		switch message.Type {
		case messagetype.Challenge:
			s.handleChallengeResponse(client, message.Payload)
		case messagetype.GetQuote:
			s.handleGetQuote(client)
		default:
			log.Printf("unknown message type: %v", message.Type)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("scanner.Err: %v", err)
	} else {
		log.Printf("connection was closed")
	}
}

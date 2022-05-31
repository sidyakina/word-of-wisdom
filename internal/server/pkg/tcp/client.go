package tcp

import (
	"encoding/json"
	"fmt"
	"github.com/sidyakina/word-of-wisdom/internal/general/challenge"
	"github.com/sidyakina/word-of-wisdom/pkg/server/api"
	"log"
	"net"
)

type Client struct {
	conn             net.Conn
	currentChallenge *challenge.Challenge
}

func newClient(conn net.Conn) *Client {
	return &Client{conn: conn}
}

func (c *Client) sendMessage(messageType string, rawPayload any) error {
	fmt.Printf("sending message with type %v and payload %+v", messageType, rawPayload)

	payload, err := json.Marshal(rawPayload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	message := api.Message{
		Type:    messageType,
		Payload: payload,
	}

	data, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	_, err = c.conn.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write message: %w", err)
	}

	log.Printf("message successfully sended to client")

	return nil
}

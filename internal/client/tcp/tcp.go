package tcp

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sidyakina/word-of-wisdom/pkg/server/api"
	messagetype "github.com/sidyakina/word-of-wisdom/pkg/server/message-type"
	"log"
	"net"
	"time"
)

type Client struct {
	conn     net.Conn
	messages chan []byte

	timeout time.Duration
}

func GetQuotes(serverAddress string, timeout time.Duration, numberQuotes int) {
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		log.Printf("failed to net.Dial: %v", err)

		return
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Printf("failed to close tcp connect: %v", err)
		}
	}()

	client := Client{
		conn:     conn,
		messages: make(chan []byte),
		timeout:  timeout,
	}

	go client.readMessages()

	for i := 1; i <= numberQuotes; i++ {
		log.Printf("request quote number %v", i)

		err = client.sendQetQuote()
		if err != nil {
			log.Printf("failed to send qet quote request: %v", err)

			return
		}

		err = client.waitAndHandle(messagetype.Challenge, client.handleChallengeRequest)
		if err != nil {
			log.Printf("challenge failed: %v", err)

			return
		}

		err = client.waitAndHandle(messagetype.GetQuote, client.handleGetQuoteResponse)
		if err != nil {
			log.Printf("challenge failed: %v", err)

			return
		}
	}
}

func (c *Client) readMessages() {
	scanner := bufio.NewScanner(c.conn)

	for scanner.Scan() {
		c.messages <- scanner.Bytes()
	}

	if err := scanner.Err(); err != nil && !errors.Is(err, net.ErrClosed) {
		log.Printf("scanner.Err: %v", err)
	} else {
		log.Printf("connection was closed")
	}
}

func (c *Client) waitAndHandle(messageType string, handle func(payload json.RawMessage) error) error {
	t := time.NewTimer(c.timeout)
	defer t.Stop()

	var rawMessage []byte
	select {
	case rawMessage = <-c.messages:
		// do nothing here
	case <-t.C:
		return fmt.Errorf("message is not received (timeout %v)", c.timeout)
	}

	message := api.Message{}
	err := json.Unmarshal(rawMessage, &message)
	if err != nil {
		return fmt.Errorf("failed to unmarshal request %s: %w", message, err)
	}

	log.Printf("new message with type %v and payload %s", message.Type, message.Payload)

	if message.Type != messageType {
		return fmt.Errorf("unexpected message type %v: expected %v", message.Type, messageType)
	}

	return handle(message.Payload)
}

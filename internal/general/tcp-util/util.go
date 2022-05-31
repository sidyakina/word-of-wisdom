package tcputil

import (
	"encoding/json"
	"fmt"
	"github.com/sidyakina/word-of-wisdom/pkg/server/api"
	"log"
	"net"
)

func SendMessage(conn net.Conn, messageType string, rawPayload any) error {
	fmt.Printf("sending message with type %v and payload %+v", messageType, rawPayload)

	var payload []byte
	var err error

	if rawPayload != nil {
		payload, err = json.Marshal(rawPayload)
		if err != nil {
			return fmt.Errorf("failed to marshal payload: %w", err)
		}
	}

	message := api.Message{
		Type:    messageType,
		Payload: payload,
	}

	data, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	_, err = conn.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write message: %w", err)
	}

	log.Printf("message successfully sended")

	return nil
}

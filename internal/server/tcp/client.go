package tcp

import (
	"encoding/json"
	"github.com/sidyakina/word-of-wisdom/internal/general/challenge"
	tcputil "github.com/sidyakina/word-of-wisdom/internal/general/tcp-util"
	"net"
)

type Client struct {
	conn             net.Conn
	currentChallenge *challenge.Challenge
}

func newClient(conn net.Conn) *Client {
	return &Client{conn: conn}
}

func (c *Client) sendMessage(messageType string, rawPayload json.Marshaler) error {
	return tcputil.SendMessage(c.conn, messageType, rawPayload)
}

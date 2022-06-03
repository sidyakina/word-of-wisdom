package tcp

import (
	tcputil "github.com/sidyakina/word-of-wisdom/internal/general/tcp-util"
	messagetype "github.com/sidyakina/word-of-wisdom/pkg/server/message-type"
)

func (c *Client) sendGetQuote() error {
	return tcputil.SendMessage(c.conn, messagetype.GetQuote, nil)
}

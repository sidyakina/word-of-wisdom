package tcp

import (
	"encoding/json"
	"fmt"
	"github.com/sidyakina/word-of-wisdom/internal/general/challenge"
	tcputil "github.com/sidyakina/word-of-wisdom/internal/general/tcp-util"
	"github.com/sidyakina/word-of-wisdom/pkg/server/api"
	messagetype "github.com/sidyakina/word-of-wisdom/pkg/server/message-type"
	"log"
)

func (c *Client) handleChallengeRequest(requestPayload json.RawMessage) error {
	payload := api.ChallengeRequestPayload{}

	err := json.Unmarshal(requestPayload, &payload)
	if err != nil {
		return fmt.Errorf("failed to unmarshal challenge request %s: %w", requestPayload, err)
	}

	currentChallenge := challenge.NewChallenge(payload.Challenge)
	log.Printf("current challenge %+v", currentChallenge)

	answer, err := currentChallenge.FindAnswer()
	if err != nil {
		return fmt.Errorf("challenge failed: %w", err)
	}

	log.Printf("answer for challenge: %v", answer)

	responsePayload := api.ChallengeResponsePayload{Answer: answer}

	return tcputil.SendMessage(c.conn, messagetype.Challenge, responsePayload)
}

func (c *Client) handleGetQuoteResponse(responsePayload json.RawMessage) error {
	payload := api.GetQuoteResponse{}

	err := json.Unmarshal(responsePayload, &payload)
	if err != nil {
		return fmt.Errorf("failed to unmarshal get quote response %s: %w", responsePayload, err)
	}

	log.Printf("received quote: %v", payload.Quote)

	return nil
}

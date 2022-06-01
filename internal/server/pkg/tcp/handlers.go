package tcp

import (
	"encoding/json"
	"github.com/sidyakina/word-of-wisdom/internal/general/challenge"
	"github.com/sidyakina/word-of-wisdom/pkg/server/api"
	messagetype "github.com/sidyakina/word-of-wisdom/pkg/server/message-type"
	"log"
)

func (s *Server) handleChallengeResponse(client *Client, responsePayload json.RawMessage) {
	if client.currentChallenge == nil {
		log.Printf("client hasn't any current challenge")

		return
	}

	response := api.ChallengeResponsePayload{}
	err := json.Unmarshal(responsePayload, &response)
	if err != nil {
		log.Printf("failed to unmarshal challenge response payload %s: %v", responsePayload, err)

		return
	}

	if !client.currentChallenge.CheckAnswer(response.Answer) {
		log.Printf("chanllenge %+v: answer %+v is wrong", client.currentChallenge, response.Answer)

		client.currentChallenge = nil // challenge is failed

		return
	}

	log.Printf("chanllenge %+v: answer %+v is right", client.currentChallenge, response.Answer)

	client.currentChallenge = nil

	s.sendQuote(client)
}

func (s *Server) sendQuote(client *Client) {
	quote, err := s.quotesRepo.GetQuote()
	if err != nil {
		log.Printf("failed to get quote: %v", err)

		return
	}

	log.Printf("quote for client: %v", quote)

	payload := api.GetQuoteResponse{Quote: quote}

	err = client.sendMessage(messagetype.GetQuote, payload)
	if err != nil {
		log.Printf("failed to send quote: %v", err)
	}
}

func (s *Server) handleGetQuote(client *Client) {
	if client.currentChallenge != nil {
		log.Printf("client already received challenge")

		return
	}

	currentChallenge := challenge.GenerateNewChallenge()
	client.currentChallenge = &currentChallenge

	log.Printf("challenge for client: %+v", currentChallenge)

	payload := api.ChallengeRequestPayload{Challenge: currentChallenge.Challenge}

	err := client.sendMessage(messagetype.Challenge, payload)
	if err != nil {
		log.Printf("failed to send challenge: %v", err)
	}
}

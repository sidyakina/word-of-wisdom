package api

type ChallengeRequestPayload struct {
	Challenge string `json:"challenge"`
}

type ChallengeResponsePayload struct {
	Answer string `json:"answer"`
}

package api

import "encoding/json"

type Message struct {
	Type    string
	Payload json.RawMessage
}

package dto

import "encoding/json"

type EventEnvelope struct {
	ID        string          `json:"id"`
	Event     string          `json:"event"`
	Data      json.RawMessage `json:"data"`
	CreatedAt int64           `json:"created_at"`
}

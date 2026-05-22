package dto

import "encoding/json"

type EventEnvelope struct {
	Event string          `json:"event"`
	Data  json.RawMessage `json:"data"`
}

package contract

import "time"

type Event struct {
	ID     uint       `json:"id"`
	Time   *time.Time `json:"time"`
	Status *string    `json:"status"`
	Data   *string    `json:"data"`
	Type   *string    `json:"type"`
}
type EventResponse struct {
	Meta *Metadata `json:"metadata"`
	Data *Event    `json:"data"`
}

type MultiEventResponse struct {
	Meta *Metadata `json:"metadata"`
	Data []Event   `json:"data"`
}

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

func (e *Event) Validate() *Error {
	if e == nil {
		return ErrBadParam("empty body")
	}
	if e.Data == nil {
		return ErrBadParam("empty param data")
	}
	if e.Time == nil {
		return ErrBadParam("empty param time")
	}
	if e.Type == nil {
		return ErrBadParam("empty param type")
	}
	return nil
}

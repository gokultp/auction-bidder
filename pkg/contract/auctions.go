package contract

import (
	"fmt"
	"time"
)

type Auction struct {
	ID            uint       `json:"id"`
	Name          *string    `json:"name"`
	Description   *string    `json:"description"`
	StartAt       *time.Time `json:"start_at"`
	EndAt         *time.Time `json:"end_at"`
	StartPrice    *uint      `json:"start_price"`
	Status        *string    `json:"status"`
	AuctionWinner *string    `json:"auction_winner"`
	CreatedBy     *uint      `json:"created_by"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
}

type AuctionResponse struct {
	Meta *Metadata `json:"metadata"`
	Data *Auction  `json:"data"`
}

type MultiAuctionResponse struct {
	Meta *Metadata `json:"metadata"`
	Data []Auction `json:"data"`
}

func (a *Auction) Validate() *Error {
	if a == nil {
		return ErrBadParam("empty body")
	}
	if a.Name == nil {
		return ErrBadParam("empty param name")
	}
	if err := validateMaxLength("name", *a.Name, 32); err != nil {
		return err
	}
	if a.Description == nil {
		return ErrBadParam("empty param description")
	}
	if err := validateMaxLength("description", *a.Description, 128); err != nil {
		return err
	}
	if a.StartAt == nil {
		return ErrBadParam("empty param start_at")
	}
	if a.EndAt == nil {
		return ErrBadParam("empty param end_at")
	}
	if a.StartPrice == nil {
		return ErrBadParam("empty param start_price")
	}
	if a.EndAt.Sub(*a.StartAt) < time.Minute*2 {
		return ErrBadParam("auction should be running atleast for 2 minutes")
	}
	if a.EndAt.Sub(time.Now()) < time.Minute*2 {
		return ErrBadParam("auction should be running atleast for 2 minutes")
	}
	return nil
}

func validateMaxLength(field, value string, l int) *Error {
	if len(value) > l {
		return ErrBadParam(fmt.Sprintf("max legth for %s is %d", field, l))
	}
	return nil
}

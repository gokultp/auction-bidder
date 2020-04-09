package contract

import "time"

type Auction struct {
	ID            uint       `json:"id"`
	Name          *string    `json:"name"`
	Description   *string    `json:"description"`
	StartAt       *time.Time `json:"start_at"`
	EndAt         *time.Time `json:"end_at"`
	StartPrice    *uint      `json:"start_price"`
	AuctionWinner *uint      `json:"auction_winner"`
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

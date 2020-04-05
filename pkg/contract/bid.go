package contract

import "time"

type Bid struct {
	ID        uint       `json:"id"`
	Price     *uint      `json:"price"`
	UserID    *uint      `json:"user_id"`
	AuctionID *uint      `json:"auction_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type BidResponse struct {
	Meta *Metadata `json:"metadata"`
	Data *Bid      `json:"data"`
}

type MultiBidResponse struct {
	Meta *Metadata `json:"metadata"`
	Data []Bid     `json:"data"`
}

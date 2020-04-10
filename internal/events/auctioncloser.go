package events

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/gokultp/auction-bidder/pkg/clients"
	"github.com/labstack/gommon/log"
)

type AuctionCloser struct {
	auctionID *string
}

func NewAuctionCloser(data *string) *AuctionCloser {
	return &AuctionCloser{
		auctionID: data,
	}
}

func (a *AuctionCloser) Exec(ctx context.Context) error {
	if a.auctionID == nil {
		return errors.New("no auction id")
	}
	e, err := clients.CloseAuction(*a.auctionID)
	d, _ := json.Marshal(e)
	log.Info(string(d))
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

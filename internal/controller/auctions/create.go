package auctions

import (
	"context"
	"fmt"

	"github.com/gokultp/auction-bidder/internal/events"
	"github.com/gokultp/auction-bidder/internal/model"
	"github.com/gokultp/auction-bidder/pkg/clients"
	"github.com/gokultp/auction-bidder/pkg/contract"
	"github.com/labstack/gommon/log"
)

func Create(ctx context.Context, auction *contract.Auction) (*contract.AuctionResponse, *contract.Error) {
	auctionData := &model.Auction{
		Name:        auction.Name,
		Description: auction.Description,
		StartAt:     auction.StartAt,
		EndAt:       auction.EndAt,
		Status:      &model.AuctionStatusActive,
		StartPrice:  auction.StartPrice,
		CreatedBy:   auction.CreatedBy,
	}
	if err := auctionData.Create(ctx); err != nil {
		log.Error(err)
		return nil, contract.ErrInternalServerError()
	}
	evtData := fmt.Sprint(auctionData.ID)
	evt := contract.Event{
		Time: auctionData.EndAt,
		Data: &evtData,
		Type: &events.EventCloseAuction,
	}
	_, evErr := clients.CreateEvent(evt)
	if evErr != nil {
		log.Error(evErr)
		return nil, contract.ErrInternalServerError()
	}
	return auctionResponse(auctionData, 200), nil
}

func auctionResponse(a *model.Auction, httpCode int) *contract.AuctionResponse {
	success := "success"
	return &contract.AuctionResponse{
		Meta: &contract.Metadata{
			Code:   &httpCode,
			Status: &success,
		},
		Data: convertAutionModelToContract(a),
	}
}

func convertAutionModelToContract(a *model.Auction) *contract.Auction {
	auction := &contract.Auction{
		ID:          a.ID,
		Name:        a.Name,
		Description: a.Description,
		StartAt:     a.StartAt,
		EndAt:       a.EndAt,
		Status:      a.Status,
		StartPrice:  a.StartPrice,
		CreatedAt:   &a.CreatedAt,
		UpdatedAt:   &a.UpdatedAt,
		CreatedBy:   a.CreatedBy,
	}

	if a.AuctionWinner != nil {
		winnerURL := fmt.Sprintf("/v1/auctions/%d/bids/%d", a.ID, *a.AuctionWinner)
		auction.AuctionWinner = &winnerURL
	}

	return auction
}

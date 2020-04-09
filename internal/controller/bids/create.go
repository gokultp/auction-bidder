package bids

import (
	"context"
	"time"

	"github.com/gokultp/auction-bidder/internal/model"
	"github.com/gokultp/auction-bidder/pkg/contract"
	"github.com/labstack/gommon/log"
)

func Create(ctx context.Context, bid *contract.Bid) (*contract.BidResponse, *contract.Error) {
	bidData := &model.Bid{
		Price:     bid.Price,
		UserID:    bid.UserID,
		AuctionID: bid.AuctionID,
	}

	auction, err := model.GetAuctionByID(ctx, *bid.AuctionID)
	if err != nil {
		log.Error(err)
		return nil, contract.ErrInternalServerError()
	}
	now := time.Now()

	if now.Before(*auction.StartAt) {
		return nil, contract.ErrBadRequest("auction is yet to start")
	}

	if now.After(*auction.EndAt) {
		return nil, contract.ErrBadRequest("auction is over")
	}

	if *bid.Price < *auction.StartPrice {
		return nil, contract.ErrBadRequest("your bid is less than start price")
	}
	if err := bidData.Create(ctx); err != nil {
		log.Error(err)
		return nil, contract.ErrInternalServerError(err.Error())
	}
	return bidResponse(bidData, 200), nil
}

func bidResponse(b *model.Bid, httpCode int) *contract.BidResponse {
	success := "success"
	return &contract.BidResponse{
		Meta: &contract.Metadata{
			Code:   &httpCode,
			Status: &success,
		},
		Data: convertBidModelToContract(*b),
	}
}

func convertBidModelToContract(b model.Bid) *contract.Bid {
	return &contract.Bid{
		ID:        b.ID,
		Price:     b.Price,
		AuctionID: b.AuctionID,
		UserID:    b.UserID,
		CreatedAt: &b.CreatedAt,
		UpdatedAt: &b.UpdatedAt,
	}
}

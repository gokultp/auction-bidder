package auctions

import (
	"context"

	"github.com/gokultp/auction-bidder/internal/model"
	"github.com/gokultp/auction-bidder/pkg/contract"
	"github.com/labstack/gommon/log"
)

func Create(ctx context.Context, auction *contract.Auction) (*contract.AuctionResponse, *contract.Error) {
	auctionData := &model.Auction{
		Name:        auction.Name,
		Description: auction.Description,
		StartAt:     auction.StartAt,
		EndAt:       auction.EndAt,
		StartPrice:  auction.StartPrice,
		CreatedBy:   auction.CreatedBy,
	}

	if err := auctionData.Create(ctx); err != nil {
		log.Error(err)
		return nil, contract.ErrInternalServerError(err.Error())
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
	return &contract.Auction{
		ID:          a.ID,
		Name:        a.Name,
		Description: a.Description,
		StartAt:     a.StartAt,
		EndAt:       a.EndAt,
		StartPrice:  a.StartPrice,
		CreatedAt:   &a.CreatedAt,
		UpdatedAt:   &a.UpdatedAt,
		CreatedBy:   a.CreatedBy,
	}
}

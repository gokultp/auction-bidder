package auctions

import (
	"context"

	"github.com/gokultp/auction-bidder/internal/model"
	"github.com/gokultp/auction-bidder/pkg/contract"
	"github.com/labstack/gommon/log"
)

func Get(ctx context.Context, auctionID *uint) (*contract.AuctionResponse, *contract.Error) {

	auction, err := model.GetAuctionByID(ctx, *auctionID)
	if err != nil {
		log.Error(err)
		return nil, contract.ErrInternalServerError()
	}
	if auction == nil {
		return nil, contract.ErrNotFound()
	}
	return auctionResponse(auction, 200), nil
}

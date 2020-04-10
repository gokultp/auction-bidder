package auctions

import (
	"context"
	"net/http"

	"github.com/gokultp/auction-bidder/internal/model"
	"github.com/gokultp/auction-bidder/pkg/contract"
	"github.com/labstack/gommon/log"
)

// Update supports only update status to complete
func Update(ctx context.Context, a *contract.Auction) (*contract.AuctionResponse, *contract.Error) {
	auctionData, err := model.GetAuctionByID(ctx, a.ID)
	if a.Status != nil && *a.Status == model.AuctionStatusClosed {
		if err != nil {
			log.Error(err)
			return nil, contract.ErrInternalServerError()
		}
		winner, err := model.GetBestBidForAuction(ctx, a.ID)
		if err != nil {
			log.Error(err)
			return nil, contract.ErrInternalServerError()
		}
		if winner != nil {
			auctionData.AuctionWinner = &winner.ID
		}
		auctionData.Status = &model.AuctionStatusClosed
		err = auctionData.Update(ctx)
		if err != nil {
			log.Error(err)
			return nil, contract.ErrInternalServerError()
		}

	}
	return auctionResponse(auctionData, http.StatusOK), nil
}

package bids

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gokultp/auction-bidder/internal/checks/auth"
	"github.com/gokultp/auction-bidder/internal/model"
	"github.com/gokultp/auction-bidder/internal/utils"
	"github.com/gokultp/auction-bidder/pkg/contract"
	"github.com/labstack/gommon/log"
)

func Get(ctx context.Context, bidID, auctionID uint) (*contract.BidResponse, *contract.Error) {
	bid, err := model.GetBidByID(ctx, bidID, auctionID)
	if err != nil {
		log.Error(err)
		return nil, contract.ErrInternalServerError()
	}
	authenticator := ctx.Value("auth").(auth.Authenticator)
	if !authenticator.IsAdmin() && *bid.UserID != authenticator.UserID() {
		return nil, contract.ErrForbidden()
	}
	return bidResponse(bid, http.StatusOK), nil
}

func BulkGet(ctx context.Context, auctionID uint, p contract.Pagination) (*contract.MultiBidResponse, *contract.Error) {
	bids, err := model.GetBidsByAuction(ctx, auctionID, p.Limit, p.Page)
	if err != nil {
		log.Error(err)
		return nil, contract.ErrInternalServerError()
	}
	bidRes := []contract.Bid{}
	for _, bid := range bids {
		bidRes = append(bidRes, *convertBidModelToContract(bid))
	}
	return &contract.MultiBidResponse{
		Meta: utils.GetMetadata(p, fmt.Sprintf("/v1/auctions/%d/bids", auctionID)),
		Data: bidRes,
	}, nil
}

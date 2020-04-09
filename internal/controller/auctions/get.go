package auctions

import (
	"context"

	"github.com/gokultp/auction-bidder/internal/model"
	"github.com/gokultp/auction-bidder/internal/utils"
	"github.com/gokultp/auction-bidder/pkg/contract"
	"github.com/labstack/gommon/log"
)

func Get(ctx context.Context, auctionID uint) (*contract.AuctionResponse, *contract.Error) {

	auction, err := model.GetAuctionByID(ctx, auctionID)
	if err != nil {
		log.Error(err)
		return nil, contract.ErrInternalServerError()
	}
	if auction == nil {
		return nil, contract.ErrNotFound()
	}
	return auctionResponse(auction, 200), nil
}

func BulkGet(ctx context.Context, p contract.Pagination) (*contract.MultiAuctionResponse, *contract.Error) {
	auctions, err := model.GetAuctions(ctx, p.Limit, p.Page)
	if err != nil {
		log.Error(err)
		return nil, contract.ErrInternalServerError()
	}

	arrAuctions := []contract.Auction{}
	for _, a := range auctions {
		arrAuctions = append(arrAuctions, *convertAutionModelToContract(&a))
	}
	return &contract.MultiAuctionResponse{
		Meta: utils.GetMetadata(p, "/v1/auctions"),
		Data: arrAuctions,
	}, nil
}

package clients

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gokultp/auction-bidder/internal/model"
	"github.com/gokultp/auction-bidder/pkg/contract"
)

func CloseAuction(auctionID string) (*contract.AuctionResponse, *contract.Error) {
	var data contract.AuctionResponse
	payload := contract.Auction{
		Status: &model.AuctionStatusClosed,
	}

	if err := NewRequest(
		http.MethodPut,
		fmt.Sprintf("http://auctionmanager/v1/auctions/%s", auctionID),
		payload,
	).SetToken(os.Getenv("ADMIN_TOKEN")).
		Dial(
			&data,
		); err != nil {
		return nil, err
	}
	return &data, nil
}

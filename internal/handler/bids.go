package handler

import (
	"context"
	"net/http"

	"github.com/gokultp/auction-bidder/internal/checks/auth"
	"github.com/gokultp/auction-bidder/internal/controller/bids"
	"github.com/gokultp/auction-bidder/pkg/contract"
	"github.com/jinzhu/gorm"
)

type BidHandler struct {
	DB *gorm.DB
}

func (h *BidHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx, err := getContext(r, h.DB)
	if err != nil {
		handleError(w, err)
		return
	}
	switch r.Method {
	case http.MethodPost:
		h.Create(ctx, w, r)
	case http.MethodGet:
		h.Get(ctx, w, r)
	default:
		handleError(w, contract.ErrMethodNotAllowed())
	}
}

func (BidHandler) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var req contract.Bid
	if err := commonHandler(ctx, r, &req, false); err != nil {
		handleError(w, err)
		return
	}

	auctionID, err, _ := getIDsFromPath(r, "auctionId")
	if err != nil {
		handleError(w, err)
		return
	}
	userID := ctx.Value("auth").(auth.Authenticator).UserID()
	req.UserID = &userID
	req.AuctionID = &auctionID
	res, cerr := bids.Create(ctx, &req)
	if cerr != nil {
		handleError(w, cerr)
		return
	}
	jsonResponse(w, res, http.StatusOK)
}

func (BidHandler) Get(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	if err := commonHandler(ctx, r, nil, false); err != nil {
		handleError(w, err)
		return
	}
	auctionID, err, _ := getIDsFromPath(r, "auctionId")
	if err != nil {
		handleError(w, err)
		return
	}
	id, err, ok := getIDsFromPath(r, "id")
	if err != nil {
		handleError(w, err)
		return
	}
	if !ok {
		p := getPagination(r)
		res, err := bids.BulkGet(ctx, uint(auctionID), p)
		if err != nil {
			handleError(w, contract.ErrBadParam("id"))
			return
		}
		jsonResponse(w, res, http.StatusOK)
		return
	}
	res, cerr := bids.Get(ctx, id, auctionID)
	if err != nil {
		handleError(w, cerr)
		return
	}
	jsonResponse(w, res, http.StatusOK)
}

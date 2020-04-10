package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gokultp/auction-bidder/internal/checks/auth"
	"github.com/gokultp/auction-bidder/internal/controller/bids"
	"github.com/gokultp/auction-bidder/pkg/contract"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
)

type BidHandler struct {
	DB *gorm.DB
}

func (h *BidHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "db", h.DB)
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
	token := getToken(r)
	if token == "" {
		handleError(w, contract.ErrUnauthorized())
		return
	}
	authenticator := auth.NewJWTAuth(token)
	if !authenticator.Authenticate() {
		handleError(w, contract.ErrUnauthorized())
		return
	}
	ctx = context.WithValue(ctx, "auth", authenticator)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Error(err)
		handleError(w, contract.ErrBadRequest())
		return
	}
	if err := validateBid(&req); err != nil {
		handleError(w, err)
		return
	}
	strAuctionId := mux.Vars(r)["auctionId"]
	auctionID, err := strconv.ParseUint(strAuctionId, 10, 64)
	if err != nil {
		log.Error(err)
		handleError(w, contract.ErrBadParam("auction_id"))
		return
	}
	auctionIDUint := uint(auctionID)
	userID := authenticator.UserID()
	req.UserID = &userID
	req.AuctionID = &auctionIDUint
	res, cerr := bids.Create(ctx, &req)
	if cerr != nil {
		handleError(w, cerr)
		return
	}
	jsonResponse(w, res, http.StatusOK)
}

func (BidHandler) Get(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	token := getToken(r)
	if token == "" {
		handleError(w, contract.ErrUnauthorized())
		return
	}

	authenticator := auth.NewJWTAuth(token)
	if !authenticator.Authenticate() {
		handleError(w, contract.ErrUnauthorized())
		return
	}
	ctx = context.WithValue(ctx, "auth", authenticator)
	strAuctionId := mux.Vars(r)["auctionId"]
	auctionID, err := strconv.ParseUint(strAuctionId, 10, 64)
	if err != nil {
		log.Error(err)
		handleError(w, contract.ErrBadParam("auction_id"))
		return
	}
	strId := mux.Vars(r)["id"]
	if strId == "" {
		p := getPagination(r)
		res, err := bids.BulkGet(ctx, uint(auctionID), p)
		if err != nil {
			handleError(w, contract.ErrBadParam("id"))
			return
		}
		jsonResponse(w, res, http.StatusOK)
		return
	}
	id, err := strconv.ParseUint(strId, 10, 64)
	if err != nil {
		log.Error(err)
		handleError(w, contract.ErrBadParam("id"))
		return
	}
	res, cerr := bids.Get(ctx, uint(id), uint(auctionID))
	if err != nil {
		handleError(w, cerr)
		return
	}
	jsonResponse(w, res, http.StatusOK)
}

func validateBid(b *contract.Bid) *contract.Error {
	if b == nil {
		return contract.ErrBadParam("empty body")
	}
	if b.Price == nil {
		return contract.ErrBadParam("empty param name")
	}

	return nil
}

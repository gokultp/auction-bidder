package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gokultp/auction-bidder/internal/checks/auth"
	"github.com/gokultp/auction-bidder/internal/controller/auctions"
	"github.com/gokultp/auction-bidder/pkg/contract"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
)

type AuctionHandler struct {
	DB *gorm.DB
}

func (h *AuctionHandler) Handle(w http.ResponseWriter, r *http.Request) {
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

func (AuctionHandler) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var req contract.Auction
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
	if !authenticator.IsAdmin() {
		handleError(w, contract.ErrForbidden())
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Error(err)
		handleError(w, contract.ErrBadRequest())
		return
	}
	if err := validateAuction(&req); err != nil {
		handleError(w, err)
		return
	}
	userID := authenticator.UserID()
	req.CreatedBy = &userID
	res, err := auctions.Create(ctx, &req)
	if err != nil {
		handleError(w, err)
		return
	}
	jsonResponse(w, res)
}

func (AuctionHandler) Get(ctx context.Context, w http.ResponseWriter, r *http.Request) {
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
	strId := mux.Vars(r)["id"]
	if strId == "" {
		p := getPagination(r)
		res, err := auctions.BulkGet(ctx, p)
		if err != nil {
			handleError(w, contract.ErrBadParam("id"))
			return
		}
		jsonResponse(w, res)
		return
	}
	id, err := strconv.ParseUint(strId, 10, 64)
	if err != nil {
		log.Error(err)
		handleError(w, contract.ErrBadParam("id"))
		return
	}
	res, cerr := auctions.Get(ctx, uint(id))
	if err != nil {
		handleError(w, cerr)
		return
	}
	jsonResponse(w, res)
}

func validateAuction(a *contract.Auction) *contract.Error {
	if a == nil {
		return contract.ErrBadParam("empty body")
	}
	if a.Name == nil {
		return contract.ErrBadParam("empty param name")
	}
	if err := validateMaxLength("name", *a.Name, 32); err != nil {
		return err
	}
	if a.Description == nil {
		return contract.ErrBadParam("empty param description")
	}
	if err := validateMaxLength("description", *a.Description, 128); err != nil {
		return err
	}
	if a.StartAt == nil {
		return contract.ErrBadParam("empty param start_at")
	}
	if a.EndAt == nil {
		return contract.ErrBadParam("empty param end_at")
	}
	if a.StartPrice == nil {
		return contract.ErrBadParam("empty param start_price")
	}
	return nil
}

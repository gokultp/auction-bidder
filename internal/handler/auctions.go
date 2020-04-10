package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gokultp/auction-bidder/internal/checks/auth"
	"github.com/gokultp/auction-bidder/internal/controller/auctions"
	"github.com/gokultp/auction-bidder/internal/model"
	"github.com/gokultp/auction-bidder/pkg/contract"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
)

type AuctionHandler struct {
	DB *gorm.DB
}

func (h *AuctionHandler) Handle(w http.ResponseWriter, r *http.Request) {
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
	case http.MethodPut:
		h.Update(ctx, w, r)
	default:
		handleError(w, contract.ErrMethodNotAllowed("route"))
	}
}

func (AuctionHandler) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var req contract.Auction
	if err := commonHandler(ctx, r, &req, true); err != nil {
		handleError(w, err)
		return
	}
	userID := ctx.Value("auth").(auth.Authenticator).UserID()
	req.CreatedBy = &userID
	res, err := auctions.Create(ctx, &req)
	if err != nil {
		handleError(w, err)
		return
	}
	jsonResponse(w, res, http.StatusOK)
}

func (AuctionHandler) Update(ctx context.Context, w http.ResponseWriter, r *http.Request) {
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
	strId := mux.Vars(r)["id"]
	if strId == "" {
		handleError(w, contract.ErrMethodNotAllowed("no id"))
		return
	}
	id, err := strconv.ParseUint(strId, 10, 64)
	if err != nil {
		log.Error(err)
		handleError(w, contract.ErrBadParam("id"))
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Error(err)
		handleError(w, contract.ErrBadRequest())
		return
	}
	if req.Status == nil || *req.Status != model.AuctionStatusClosed {
		handleError(w, contract.ErrBadRequest("only allowed action is close auction"))
		return
	}
	req.ID = uint(id)
	res, cerr := auctions.Update(ctx, &req)
	if cerr != nil {
		handleError(w, cerr)
		return
	}
	jsonResponse(w, res, http.StatusOK)
}

func (AuctionHandler) Get(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	if err := commonHandler(ctx, r, nil, false); err != nil {
		handleError(w, contract.ErrBadParam("id"))
		return
	}
	id, err, ok := getIDsFromPath(r, "id")
	if err != nil {
		handleError(w, contract.ErrBadParam("id"))
		return
	}
	if !ok {
		p := getPagination(r)
		res, err := auctions.BulkGet(ctx, p)
		if err != nil {
			handleError(w, contract.ErrBadParam("id"))
			return
		}
		jsonResponse(w, res, http.StatusOK)
		return
	}
	res, cerr := auctions.Get(ctx, id)
	if err != nil {
		handleError(w, cerr)
		return
	}
	jsonResponse(w, res, http.StatusOK)
}

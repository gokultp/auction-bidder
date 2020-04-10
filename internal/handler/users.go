package handler

import (
	"context"
	"net/http"

	"github.com/gokultp/auction-bidder/internal/checks/auth"
	"github.com/gokultp/auction-bidder/internal/controller/users"
	"github.com/gokultp/auction-bidder/pkg/contract"
	"github.com/jinzhu/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func (h *UserHandler) Handle(w http.ResponseWriter, r *http.Request) {
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

func (UserHandler) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var req contract.User
	if err := commonHandler(ctx, r, &req, true); err != nil {
		handleError(w, err)
		return
	}

	res, err := users.Create(ctx, &req)
	if err != nil {
		handleError(w, err)
		return
	}
	jsonResponse(w, res, http.StatusOK)
}

func (UserHandler) Get(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	id, err, ok := getIDsFromPath(r, "id")
	if err != nil {
		handleError(w, err)
		return
	}
	if !ok {
		handleError(w, contract.ErrMethodNotAllowed())
		return
	}
	if err := commonHandler(ctx, r, nil, false); err != nil {
		handleError(w, err)
		return
	}
	if a := ctx.Value("auth").(auth.Authenticator); !a.IsAdmin() && a.UserID() != uint(id) {
		handleError(w, contract.ErrForbidden())
		return
	}
	res, cerr := users.Get(ctx, uint(id))
	if err != nil {
		handleError(w, cerr)
		return
	}
	jsonResponse(w, res, http.StatusOK)
}

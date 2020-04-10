package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gokultp/auction-bidder/internal/checks/auth"
	"github.com/gokultp/auction-bidder/internal/controller/users"
	"github.com/gokultp/auction-bidder/pkg/contract"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
)

type UserHandler struct {
	DB *gorm.DB
}

func (h *UserHandler) Handle(w http.ResponseWriter, r *http.Request) {
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

func (UserHandler) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var req contract.User
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
	if err := validateUser(&req); err != nil {
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
	strId := mux.Vars(r)["id"]
	if strId == "" {
		// no bulk get for users
		handleError(w, contract.ErrMethodNotAllowed())
		return
	}
	id, err := strconv.ParseUint(strId, 10, 64)
	if err != nil {
		log.Error(err)
		handleError(w, contract.ErrBadParam("id"))
		return
	}

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
	if !authenticator.IsAdmin() && authenticator.UserID() != uint(id) {
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

func validateUser(u *contract.User) *contract.Error {
	if u == nil {
		return contract.ErrBadParam("empty body")
	}
	if u.FirstName == nil {
		return contract.ErrBadParam("empty param first_name")
	}
	if u.LastName == nil {
		return contract.ErrBadParam("empty param last_name")
	}
	if u.Email == nil {
		return contract.ErrBadParam("empty param email")
	}
	return nil
}

package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gokultp/auction-bidder/internal/checks/auth"
	"github.com/gokultp/auction-bidder/pkg/contract"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
)

type Request interface {
	Validate() *contract.Error
}

func commonHandler(ctx context.Context, r *http.Request, req Request, adminAction bool) *contract.Error {
	token := getToken(r)
	if token == "" {
		return contract.ErrUnauthorized()
	}
	authenticator := ctx.Value("auth").(auth.Authenticator)
	if authenticator != nil && !authenticator.Authenticate() {
		return contract.ErrUnauthorized()
	}
	if adminAction && !authenticator.IsAdmin() {
		return contract.ErrForbidden()
	}
	if req != nil && r.Body != nil {
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error(err)
			return contract.ErrBadRequest()
		}
		if err := req.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func getContext(r *http.Request, db *gorm.DB) (context.Context, *contract.Error) {
	ctx := context.WithValue(r.Context(), "db", db)
	token := getToken(r)
	if token == "" {
		return ctx, contract.ErrUnauthorized()
	}
	return context.WithValue(ctx, "auth", auth.NewJWTAuth(token)), nil

}

func getIDsFromPath(r *http.Request, key string) (uint, *contract.Error, bool) {
	strId := mux.Vars(r)[key]
	if strId == "" {
		// no bulk get for users
		return 0, nil, false
	}
	id, err := strconv.ParseUint(strId, 10, 64)
	if err != nil {

		return 0, contract.ErrBadParam("id"), false
	}
	return uint(id), nil, true
}

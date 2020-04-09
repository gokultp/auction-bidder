package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gokultp/auction-bidder/internal/controller/events"
	"github.com/gokultp/auction-bidder/pkg/contract"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
)

type EventHandler struct {
	DB *gorm.DB
}

func (h *EventHandler) Handle(w http.ResponseWriter, r *http.Request) {
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

func (EventHandler) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var req contract.Event

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Error(err)
		handleError(w, contract.ErrBadRequest())
		return
	}
	if err := validateEvent(&req); err != nil {
		handleError(w, err)
		return
	}
	res, err := events.Create(ctx, &req)
	if err != nil {
		handleError(w, err)
		return
	}
	jsonResponse(w, res)
}

func (EventHandler) Get(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	strId := mux.Vars(r)["id"]
	if strId == "" {
		// no bulk get for events
		handleError(w, contract.ErrMethodNotAllowed())
		return
	}
	id, err := strconv.ParseUint(strId, 10, 64)
	if err != nil {
		log.Error(err)
		handleError(w, contract.ErrBadParam("id"))
		return
	}
	res, cerr := events.Get(ctx, uint(id))
	if err != nil {
		handleError(w, cerr)
		return
	}
	jsonResponse(w, res)
}

func validateEvent(e *contract.Event) *contract.Error {
	if e == nil {
		return contract.ErrBadParam("empty body")
	}
	if e.Data == nil {
		return contract.ErrBadParam("empty param data")
	}
	if e.Time == nil {
		return contract.ErrBadParam("empty param time")
	}
	return nil
}

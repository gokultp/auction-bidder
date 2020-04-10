package handler

import (
	"context"
	"net/http"

	"github.com/gokultp/auction-bidder/internal/controller/events"
	"github.com/gokultp/auction-bidder/pkg/contract"
	"github.com/jinzhu/gorm"
	"github.com/kr/beanstalk"
)

type EventHandler struct {
	DB    *gorm.DB
	Queue *beanstalk.Conn
}

func (h *EventHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx, err := getContext(r, h.DB)
	if err != nil {
		handleError(w, err)
		return
	}
	ctx = context.WithValue(ctx, "queue", h.Queue)

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

	if err := commonHandler(ctx, r, &req, true); err != nil {
		handleError(w, err)
		return
	}
	res, err := events.Create(ctx, &req)
	if err != nil {
		handleError(w, err)
		return
	}
	jsonResponse(w, res, http.StatusOK)
}

func (EventHandler) Get(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	if err := commonHandler(ctx, r, nil, true); err != nil {
		handleError(w, err)
		return
	}
	id, err, ok := getIDsFromPath(r, "id")
	if err != nil {
		handleError(w, err)
		return
	}
	if !ok {
		handleError(w, contract.ErrMethodNotAllowed())
		return
	}
	res, cerr := events.Get(ctx, id)
	if err != nil {
		handleError(w, cerr)
		return
	}
	jsonResponse(w, res, http.StatusOK)
}

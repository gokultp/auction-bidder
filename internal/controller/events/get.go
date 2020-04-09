package events

import (
	"context"

	"github.com/gokultp/auction-bidder/internal/model"
	"github.com/gokultp/auction-bidder/pkg/contract"
	"github.com/labstack/gommon/log"
)

func Get(ctx context.Context, eventID uint) (*contract.EventResponse, *contract.Error) {

	event, err := model.GetEventByID(ctx, eventID)
	if err != nil {
		log.Error(err)
		return nil, contract.ErrInternalServerError()
	}
	if event == nil {
		return nil, contract.ErrNotFound()
	}
	return eventResponse(event, 200), nil
}

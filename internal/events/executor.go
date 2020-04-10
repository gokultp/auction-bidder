package events

import (
	"context"

	"github.com/gokultp/auction-bidder/internal/model"
	"github.com/gokultp/auction-bidder/pkg/contract"
)

var (
	EventCloseAuction = "close-auction"
)

type Executor interface {
	Exec(ctx context.Context) error
}

func NewExecutor(e contract.Event) Executor {
	switch *e.Type {
	case EventCloseAuction:
		return NewAuctionCloser(e.Data)
	default:
		return NewAuctionCloser(e.Data)
	}
}

func ExecuteEvent(ctx context.Context, e contract.Event) error {
	ex := NewExecutor(e)
	if err := ex.Exec(ctx); err != nil {
		return err
	}

	evtModel, err := model.GetEventByID(ctx, e.ID)
	if err != nil {
		return err
	}
	evtModel.Status = &model.EventStatusCompleted
	return evtModel.Update(ctx)
}

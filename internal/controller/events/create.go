package events

import (
	"context"
	"time"

	"github.com/gokultp/auction-bidder/internal/events"
	"github.com/gokultp/auction-bidder/internal/model"
	"github.com/gokultp/auction-bidder/pkg/contract"
	"github.com/kr/beanstalk"
	"github.com/labstack/gommon/log"
)

func Create(ctx context.Context, event *contract.Event) (*contract.EventResponse, *contract.Error) {
	eventData := &model.Event{
		Time:   event.Time,
		Data:   event.Data,
		Status: &model.EventStatusScheduled,
		Type:   event.Type,
	}

	if err := eventData.Create(ctx); err != nil {
		log.Error(err)
		return nil, contract.ErrInternalServerError(err.Error())
	}
	if eventData.Time.Sub(time.Now()) < 5*time.Minute {
		qconn := ctx.Value("queue").(*beanstalk.Conn)
		events.EnqueueEvent(ctx, eventData, qconn)
	}

	return eventResponse(eventData, 200), nil
}

func eventResponse(e *model.Event, httpCode int) *contract.EventResponse {
	success := "success"
	return &contract.EventResponse{
		Meta: &contract.Metadata{
			Code:   &httpCode,
			Status: &success,
		},
		Data: ConvertEventToContract(e),
	}
}

func ConvertEventToContract(a *model.Event) *contract.Event {
	return &contract.Event{
		ID:     a.ID,
		Time:   a.Time,
		Status: a.Status,
		Data:   a.Data,
		Type:   a.Type,
	}
}

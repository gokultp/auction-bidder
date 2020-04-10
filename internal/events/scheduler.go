package events

import (
	"context"
	"encoding/json"
	"time"

	"github.com/gokultp/auction-bidder/internal/model"
	"github.com/gokultp/auction-bidder/pkg/contract"
	"github.com/kr/beanstalk"
	"github.com/labstack/gommon/log"
)

func EnqueueEvent(ctx context.Context, e *model.Event, conn *beanstalk.Conn) error {
	event := contract.Event{
		ID:     e.ID,
		Time:   e.Time,
		Data:   e.Data,
		Status: e.Status,
		Type:   e.Type,
	}
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}
	id, err := conn.Put(data, 1, max(0, e.Time.Sub(time.Now())), 120*time.Second*5)
	if err != nil {
		return err
	}
	log.Info("enqueued event", event.ID, "job id", id)
	e.Status = &model.EventStatusQueued
	err = e.Update(ctx)
	if err != nil {
		return err
	}
	return nil
}

func max(a, b time.Duration) time.Duration {
	if a > b {
		return a
	}
	return b
}

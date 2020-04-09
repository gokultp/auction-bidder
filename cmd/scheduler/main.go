package main

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/gokultp/auction-bidder/internal/db"
	"github.com/gokultp/auction-bidder/internal/model"
	"github.com/kr/beanstalk"
	"github.com/labstack/gommon/log"
)

func main() {
	dbConn, err := db.InitDB()
	if err != nil {
		panic(err)
	}

	ctx := context.WithValue(context.Background(), "db", dbConn)

	qconn, err := beanstalk.Dial("tcp", "beanstalk:11300")
	if err != nil {
		panic(err)
	}
	scheduleEvents(ctx, qconn)
}

func scheduleEvents(ctx context.Context, qconn *beanstalk.Conn) {
	events, err := model.GetEventsForNextBatch(ctx)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	for _, event := range events {
		wg.Add(1)
		go func(e model.Event, wg *sync.WaitGroup) {
			defer wg.Done()
			data, err := json.Marshal(e)
			if err != nil {
				log.Error(err)
				return
			}
			id, err := qconn.Put(data, 1, 0, max(0, e.Time.Sub(time.Now())))
			if err != nil {
				log.Error(err)
				return
			}
			e.Status = &model.EventStatusScheduled
			err = e.Update(ctx)
			if err != nil {
				log.Error(err)
				return
			}
			log.Info("enqueued event", e.ID, "job id", id)
		}(event, &wg)
	}
	wg.Wait()
}

func max(a, b time.Duration) time.Duration {
	if a > b {
		return a
	}
	return b
}

package main

import (
	"context"
	"sync"

	"github.com/gokultp/auction-bidder/internal/db"
	"github.com/gokultp/auction-bidder/internal/events"
	"github.com/gokultp/auction-bidder/internal/model"
	"github.com/kr/beanstalk"
)

func main() {
	dbConn, err := db.InitDB()
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()
	ctx := context.WithValue(context.Background(), "db", dbConn)

	qconn, err := beanstalk.Dial("tcp", "beanstalk:11300")
	if err != nil {
		panic(err)
	}
	defer qconn.Close()
	scheduleEvents(ctx, qconn)
}

func scheduleEvents(ctx context.Context, qconn *beanstalk.Conn) {
	eventsForNextBatch, err := model.GetEventsForNextBatch(ctx)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	for _, event := range eventsForNextBatch {
		wg.Add(1)
		go func(e model.Event, wg *sync.WaitGroup) {
			defer wg.Done()
			events.EnqueueEvent(ctx, &e, qconn)
		}(event, &wg)
	}
	wg.Wait()
}

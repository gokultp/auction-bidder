package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/gokultp/auction-bidder/internal/db"
	"github.com/gokultp/auction-bidder/internal/events"
	"github.com/gokultp/auction-bidder/pkg/contract"
	"github.com/kr/beanstalk"
	"github.com/labstack/gommon/log"
)

func main() {
	numberOfWorkers := 5
	log.Info("starting workers")

	for i := 0; i < numberOfWorkers; i++ {
		go startWorker()
	}
	select {}
}

func startWorker() {
	c, err := beanstalk.Dial("tcp", "beanstalk:11300")
	if err != nil {
		log.Error(err)
		return
	}
	defer c.Close()
	dbConn, err := db.InitDB()
	if err != nil {
		log.Error(err)
		return
	}
	defer dbConn.Close()
	ctx := context.WithValue(context.Background(), "db", dbConn)
	for {
		id, body, err := c.Reserve(5 * time.Second)
		if err != nil {
			continue
		}
		err = processJob(ctx, body)
		if err != nil {
			log.Error(err)
		}
		c.Delete(id)
	}
}

func processJob(ctx context.Context, data []byte) error {
	var evt contract.Event
	log.Info("got job", string(data))
	if err := json.Unmarshal(data, &evt); err != nil {
		return err
	}
	if evt.Type == nil {
		return nil
	}

	return events.ExecuteEvent(ctx, evt)
}

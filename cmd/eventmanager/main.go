package main

import (
	"flag"
	"net/http"

	"github.com/gokultp/auction-bidder/internal/db"
	"github.com/gokultp/auction-bidder/internal/handler"
	"github.com/gokultp/auction-bidder/internal/utils"
	"github.com/gokultp/auction-bidder/pkg/uptime"
	"github.com/gorilla/mux"
	"github.com/kr/beanstalk"
	"github.com/labstack/gommon/log"
)

var (
	// Version is the build version
	Version string
	// MinVersion is the latest git commit hash
	MinVersion string
	// BuildTime is the time at which this build is made
	BuildTime string
	port      string
)

func main() {
	flag.StringVar(&port, "port", "80", "--port 80")
	dbConn, err := db.InitDB()
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queueConn, err := beanstalk.Dial("tcp", "beanstalk:11300")
	if err != nil {
		panic(err)
	}

	err = utils.InitJWT("./id_rsa", "./id_rsa.pub")
	if err != nil {
		panic(err)
	}

	defer queueConn.Close()

	u := uptime.NewUptime(Version, MinVersion, BuildTime)
	r := mux.NewRouter()
	r.NotFoundHandler = handler.NotFoundHandler{}
	e := &handler.EventHandler{
		DB:    dbConn,
		Queue: queueConn,
	}
	r.HandleFunc("/health", u.Handler)
	r.HandleFunc("/v1/events", e.Handle)
	r.HandleFunc("/v1/events/{id:[0-9]+}", e.Handle)
	log.Info("listening at", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Error(err)
	}
}

package main

import (
	"flag"
	"net/http"

	"github.com/gokultp/auction-bidder/internal/db"
	"github.com/gokultp/auction-bidder/internal/handler"
	"github.com/gokultp/auction-bidder/pkg/uptime"
	"github.com/gorilla/mux"
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

	u := uptime.NewUptime(Version, MinVersion, BuildTime)
	r := mux.NewRouter()
	r.NotFoundHandler = handler.NotFoundHandler{}
	e := &handler.EventHandler{DB: dbConn}
	r.HandleFunc("/health", u.Handler)
	r.HandleFunc("/v1/events", e.Handle)
	r.HandleFunc("/v1/events/{id:[0-9]+}", e.Handle)

	http.ListenAndServe(":"+port, r)
}

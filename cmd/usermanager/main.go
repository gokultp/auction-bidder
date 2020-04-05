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
	db, err := db.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	u := uptime.NewUptime(Version, MinVersion, BuildTime)
	r := mux.NewRouter()
	users := &handler.UserHandler{DB: db}
	r.HandleFunc("/health", u.Handler)
	r.HandleFunc("/v1/users", users.Handle)
	r.HandleFunc("/v1/users/{id:[0-9]+}", users.Handle)

	http.ListenAndServe(":"+port, r)
}
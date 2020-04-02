package main

import (
	"flag"
	"net/http"

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
	u := uptime.NewUptime(Version, MinVersion, BuildTime)
	r := mux.NewRouter()
	r.HandleFunc("/health", u.Handler)

	http.ListenAndServe(":"+port, r)
}

package main

import (
	"flag"
	"net/http"

	"github.com/gokultp/auction-bidder/internal/db"
	"github.com/gokultp/auction-bidder/internal/handler"
	"github.com/gokultp/auction-bidder/internal/utils"
	"github.com/gokultp/auction-bidder/pkg/uptime"
	"github.com/gorilla/mux"
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

	err = utils.InitJWT("./id_rsa", "./id_rsa.pub")
	if err != nil {
		panic(err)
	}

	if err := db.InitAdmin(dbConn); err != nil {
		panic(err)
	}

	u := uptime.NewUptime(Version, MinVersion, BuildTime)
	r := mux.NewRouter()
	r.NotFoundHandler = handler.NotFoundHandler{}
	users := &handler.UserHandler{DB: dbConn}
	r.HandleFunc("/health", u.Handler)
	r.HandleFunc("/v1/users", users.Handle)
	r.HandleFunc("/v1/users/{id:[0-9]+}", users.Handle)
	log.Info("Listening at ", port)
	http.ListenAndServe(":"+port, r)
}

package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gokultp/auction-bidder/internal/db"
	"github.com/gokultp/auction-bidder/internal/handler"
	"github.com/gokultp/auction-bidder/internal/utils"
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
	err = utils.InitJWT("./id_rsa", "./id_rsa.pub")
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.NotFoundHandler = handler.NotFoundHandler{}
	u := uptime.NewUptime(Version, MinVersion, BuildTime)
	r.HandleFunc("/health", u.Handler)

	auctions := &handler.AuctionHandler{DB: db}
	r.HandleFunc("/v1/auctions", auctions.Handle)
	r.HandleFunc("/v1/auctions/{id:[0-9]+}", auctions.Handle)

	bids := &handler.BidHandler{DB: db}
	r.HandleFunc("/v1/auctions/{auctionId:[0-9]+}/bids", bids.Handle)
	r.HandleFunc("/v1/auctions/{auctionId:[0-9]+}/bids/{id:[0-9]+}", bids.Handle)

	fmt.Println("listening on port ", port)
	http.ListenAndServe(":"+port, r)
}

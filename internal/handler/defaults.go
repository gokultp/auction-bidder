package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gokultp/auction-bidder/pkg/contract"
	"github.com/labstack/gommon/log"
)

func handleError(w http.ResponseWriter, err *contract.Error) {
	jsonResponse(w, contract.NewErrorResponse(err))
}

func jsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Error(err)
	}
}

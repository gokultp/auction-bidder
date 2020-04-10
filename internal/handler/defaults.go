package handler

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gokultp/auction-bidder/pkg/contract"
	"github.com/labstack/gommon/log"
)

type NotFoundHandler struct{}

func (NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handleError(w, contract.ErrNotFound())
}

func handleError(w http.ResponseWriter, err *contract.Error) {
	jsonResponse(w, contract.NewErrorResponse(err), err.HTTPCode)
}

func jsonResponse(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Error(err)
	}
}

func getToken(r *http.Request) string {
	var re = regexp.MustCompile(`(?mi)bearer `)
	return re.ReplaceAllString(r.Header.Get("authorization"), "")
}

func getPagination(r *http.Request) contract.Pagination {
	p := contract.Pagination{
		Page:  1,
		Limit: 10,
	}
	query := r.URL.Query()
	strPage := query.Get("page")
	if strPage != "" {
		if page, err := strconv.ParseUint(strPage, 10, 64); err == nil {
			p.Page = uint(page)
		}
	}
	strLimit := query.Get("limit")
	if strLimit != "" {
		if limit, err := strconv.ParseUint(strLimit, 10, 64); err == nil {
			p.Limit = uint(limit)
		}
	}
	return p
}

package utils

import (
	"fmt"
	"net/url"

	"github.com/gokultp/auction-bidder/pkg/contract"
)

func GetMetadata(p contract.Pagination, uri string) *contract.Metadata {
	statusCodeOK := 200
	statusSuccess := "success"
	meta := contract.Metadata{
		NextPage: getPageURL(uri, p.Page+1, p.Limit),
		Page:     &p.Page,
		Limit:    &p.Limit,
		Code:     &statusCodeOK,
		Status:   &statusSuccess,
	}
	if p.Page > 1 {
		meta.PrevPage = getPageURL(uri, p.Page-1, p.Limit)
	}
	return &meta
}

func getPageURL(uri string, page, limit uint) *string {
	u, _ := url.Parse(uri)
	q := u.Query()
	q.Set("page", fmt.Sprint(page))
	q.Set("limit", fmt.Sprint(limit))
	u.RawQuery = q.Encode()
	uri = u.String()
	return &uri
}

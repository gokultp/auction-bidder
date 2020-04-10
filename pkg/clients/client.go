package clients

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gokultp/auction-bidder/pkg/contract"
)

type Request struct {
	url    string
	body   interface{}
	method string
	http.Header
}

func NewRequest(method, url string, body interface{}) *Request {
	return &Request{
		method: method,
		url:    url,
		body:   body,
	}
}
func (r *Request) SetToken(token string) *Request {
	if r.Header == nil {
		r.Header = make(http.Header)
	}
	r.Header.Set("authorization", "bearer "+token)
	return r
}
func (r *Request) Dial(response interface{}) *contract.Error {
	payload, err := json.Marshal(r.body)
	if err != nil {
		return contract.ErrBadRequest(err.Error())
	}
	req, err := http.NewRequest(r.method, r.url, bytes.NewReader(payload))
	if err != nil {
		return contract.ErrBadRequest(err.Error())
	}
	if r.Header != nil {
		req.Header = r.Header
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return contract.ErrBadRequest(err.Error())
	}
	if res.StatusCode != http.StatusOK {
		var errData contract.ErrorResponse
		if err := json.NewDecoder(res.Body).Decode(&errData); err != nil {
			return contract.ErrBadRequest(err.Error())
		}
		return errData.Error
	}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return contract.ErrBadRequest(err.Error())
	}
	return nil
}

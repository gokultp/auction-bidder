package clients

import (
	"net/http"
	"os"

	"github.com/gokultp/auction-bidder/pkg/contract"
)

func CreateEvent(evt contract.Event) (*contract.EventResponse, *contract.Error) {
	var data contract.EventResponse

	if err := NewRequest(http.MethodPost, "http://eventmanager/v1/events", evt).
		SetToken(os.Getenv("ADMIN_TOKEN")).
		Dial(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

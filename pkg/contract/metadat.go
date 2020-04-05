package contract

var (
	StatusSuccess = "success"
	StatusFailed  = "failed"
)

type Metadata struct {
	NextPage *string `json:"next_page,omitempty"`
	PrevPage *string `json:"prev_page,omitempty"`
	Page     *uint   `json:"page,omitempty"`
	Limit    *uint   `json:"limit,omitempty"`
	Code     *int    `json:"code,omitempty"`
	Status   *string `json:"status,omitempty"`
}

type ErrorResponse struct {
	Meta  *Metadata `json:"metadata"`
	Error *Error    `json:"error"`
}

func NewErrorResponse(err *Error) *ErrorResponse {
	return &ErrorResponse{
		Meta: &Metadata{
			Status: &StatusFailed,
			Code:   &err.HTTPCode,
		},
		Error: err,
	}
}

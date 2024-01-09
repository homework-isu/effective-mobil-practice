package response


import (
	"github.com/pkg/errors"
)


type BadResponse struct {
	Error string `json:"error"`
}

func NewBadResponse(err error) *BadResponse {
	new_err := errors.Unwrap(err)
	if new_err != nil {
		return &BadResponse{
			Error: new_err.Error(),
		}
	}
	return &BadResponse{
		Error: err.Error(),
	}
}
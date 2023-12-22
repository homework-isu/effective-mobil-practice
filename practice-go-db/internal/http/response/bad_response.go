package response


import (
	"github.com/pkg/errors"
)


type BadResponse struct {
	Error string `json:"error"`
}

func NewBadResponse(err error) *BadResponse {
	return &BadResponse{
		Error: errors.Unwrap(err).Error(),
	}
}
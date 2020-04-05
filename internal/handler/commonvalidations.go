package handler

import (
	"fmt"

	"github.com/gokultp/auction-bidder/pkg/contract"
)

func validateMaxLength(field, value string, l int) *contract.Error {
	if len(value) > l {
		return contract.ErrBadParam(fmt.Sprintf("max legth for %s is %d", field, l))
	}
	return nil
}

package users

import (
	"context"

	"github.com/gokultp/auction-bidder/internal/model"
	"github.com/gokultp/auction-bidder/pkg/contract"
	"github.com/labstack/gommon/log"
)

func Get(ctx context.Context, userID uint) (*contract.UserResponse, *contract.Error) {
	user, err := model.GetUserByID(ctx, userID)
	if err != nil {
		log.Error(err)
		return nil, contract.ErrInternalServerError()
	}
	return userResponse(user), nil
}

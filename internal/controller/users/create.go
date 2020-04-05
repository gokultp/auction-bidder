package users

import (
	"context"
	"fmt"

	"github.com/gokultp/auction-bidder/internal/model"
	"github.com/gokultp/auction-bidder/pkg/contract"
	"github.com/labstack/gommon/log"
)

var (
	varTrue  = true
	varFalse = false
)

func Create(ctx context.Context, user *contract.User) (*contract.UserResponse, *contract.Error) {
	userData := &model.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		IsAdmin:   user.IsAdmin,
		IsActive:  &varTrue,
	}
	if userData.IsAdmin == nil {
		userData.IsAdmin = &varFalse
	}

	userByEmail, err := model.GetUserByEmail(ctx, *userData.Email)
	if err != nil {
		log.Error(err)
		return nil, contract.ErrInternalServerError(err.Error())
	}

	if userByEmail != nil {
		return nil, contract.ErrBadParam(fmt.Sprintf("another user exist with email %s", *userByEmail.Email))
	}

	if err := userData.Create(ctx); err != nil {
		log.Error(err)
		return nil, contract.ErrInternalServerError(err.Error())
	}
	return userResponse(userData), nil
}

func userResponse(user *model.User) *contract.UserResponse {
	success := "success"
	statusOK := 200
	return &contract.UserResponse{
		Meta: &contract.Metadata{
			Code:   &statusOK,
			Status: &success,
		},
		Data: &contract.User{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Token:     user.Token,
			IsAdmin:   user.IsAdmin,
			IsActive:  user.IsActive,
			CreatedAt: &user.CreatedAt,
			UpdatedAt: &user.UpdatedAt,
		},
	}
}

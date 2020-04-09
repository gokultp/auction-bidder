package db

import (
	"context"

	"github.com/gokultp/auction-bidder/internal/model"
	"github.com/jinzhu/gorm"
)

func InitAdmin(db *gorm.DB) error {
	ctx := context.WithValue(context.Background(), "db", db)
	user, err := model.GetUserByID(ctx, 1)
	if err != nil {
		return err
	}
	fname := "Gokul"
	email := "tp.gokul@gmail.com"
	isAdmin := true
	isActive := true
	token := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjUxODY0NDcwODAsImlhdCI6MTU4NjQ0NzA4MCwiaXNzIjoidXNlcm1hbmFnZXIiLCJuYmYiOjE1ODY0NDcwODAsInVzZXJfaWQiOjEsImlzX2FkbWluIjp0cnVlfQ.URTPH2f6Wwro-uNXT43ooMpZyHfUpf3LnouaAadAEX4tkm9Cao9Uui7LO_ufIP47S_XxHKiZ1w7YRrkZvnGChM6BhAc6Dy6hzA12qM_V8SUpl_oasiIu1gQzfZQhB7IJs2f_NW28mj1ulpYyd3_hPU27EX6WYb-ftElK9t9i9gW1mXWbt3tDJ6u3-bKaPJ9tjNrJgpFHtFWx48C6rgm4h4GwF_FxWbTulWr91XK3dhYk5yLSExseAT6yQ9CnSJUluMzHyNrg5U9YQMUjQGGfYyNxfq8DFIQWC4J9b8HgNtSiiDqo8FLv80drQzyyvkSda1VUxXepaC1d9gpgsCsOXA"
	if user == nil {
		user := model.User{
			FirstName: &fname,
			IsAdmin:   &isAdmin,
			IsActive:  &isActive,
			Token:     &token,
			Email:     &email,
		}
		return user.Create(ctx)
	}
	return nil
}

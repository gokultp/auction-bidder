package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gokultp/auction-bidder/internal/utils"
	"github.com/gokultp/auction-bidder/pkg/contract"
	"github.com/labstack/gommon/log"
)

type JWTAuth struct {
	token       string
	accessToken *contract.AccessToken
}

func NewJWTAuth(token string) *JWTAuth {
	return &JWTAuth{
		token: token,
	}
}

func (a *JWTAuth) Authenticate() bool {
	token, err := jwt.ParseWithClaims(a.token, &contract.AccessToken{}, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return utils.PublicKey, nil
	})
	if err != nil {
		log.Error(err)
		return false
	}

	if err := token.Claims.Valid(); err != nil {
		log.Error(err)
		return false
	}
	a.accessToken = token.Claims.(*contract.AccessToken)
	return true
}

func (a JWTAuth) UserID() uint {
	return a.accessToken.UserID
}

func (a JWTAuth) IsAdmin() bool {
	return a.accessToken.IsAdmin
}

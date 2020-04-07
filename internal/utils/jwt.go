package utils

import (
	"crypto/rsa"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gokultp/auction-bidder/pkg/contract"
	"github.com/labstack/gommon/log"
)

var (
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
)

func InitJWT(pemFile, pubFile string) error {
	pem, err := ioutil.ReadFile(pemFile)
	if err != nil {
		log.Error(err)
		return err
	}
	PrivateKey, err = jwt.ParseRSAPrivateKeyFromPEM(pem)
	if err != nil {
		log.Error(err)
		return err
	}
	pub, err := ioutil.ReadFile(pubFile)
	if err != nil {
		log.Error(err)
		return err
	}
	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(pub)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func GenerateToken(userid uint, isAdmin bool) (string, error) {
	claims := contract.AccessToken{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1000000 * time.Hour).Unix(),
			Issuer:    "usermanager",
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
		},
		contract.TokenBody{
			UserID:  userid,
			IsAdmin: isAdmin,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(PrivateKey)
}

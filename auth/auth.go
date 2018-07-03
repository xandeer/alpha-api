package auth

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"

	. "github.com/xandeer/alpha-api/config"
)

var config = Config{}
var (
	VerifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

type CustomerInfo struct {
	Name string
	Kind string
}

type Claims struct {
	*jwt.StandardClaims
	TokenType string
	CustomerInfo
}

func init() {
	config.Read()

	signBytes, err := ioutil.ReadFile(config.PrivKeyPath)
	fatal(err)

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	fatal(err)

	verifyBytes, err := ioutil.ReadFile(config.PubKeyPath)
	fatal(err)

	VerifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	fatal(err)
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CreateToken(user string) (string, error) {
	t := jwt.New(jwt.GetSigningMethod("RS256"))

	t.Claims = &Claims{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
		},
		"level1",
		CustomerInfo{user, "admin"},
	}

	return t.SignedString(signKey)
}

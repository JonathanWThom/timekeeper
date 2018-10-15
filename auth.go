package main

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	privKeyPath = "keys/private.key"
	pubKeyPath  = "keys/public.key.pub"
)

var verifyKey *rsa.PublicKey
var signKey *rsa.PrivateKey

func (s *server) initKeys() {
	var err error
	var signKeyByte []byte
	if os.Getenv("PORT") == "" {
		signKeyByte, err = ioutil.ReadFile(privKeyPath)
	} else {
		signKeyByte = []byte(os.Getenv("PRIVATE_KEY"))
	}
	if err != nil {
		log.Fatalf("[privateKey]: %s\n", err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signKeyByte)
	if err != nil {
		log.Fatalf("[privateKey]: %s\n", err)
	}

	var verifyKeyByte []byte
	if os.Getenv("PORT") == "" {
		verifyKeyByte, err = ioutil.ReadFile(pubKeyPath)
	} else {
		verifyKeyByte = []byte(os.Getenv("PUBLIC_KEY"))
	}
	if err != nil {
		log.Fatalf("[publicKey]: %s\n", err)
		panic(err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyKeyByte)
	if err != nil {
		log.Fatalf("[publicKey]: %s\n", err)
		panic(err)
	}
}

func (s *server) validateTokenMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
			func(token *jwt.Token) (interface{}, error) {
				return verifyKey, nil
			})

		if err != nil {
			jsonUnauthorized(err, w, r)
			return
		}

		if token.Valid {
			next.ServeHTTP(w, r)
		} else {
			jsonUnauthorized(err, w, r)
			return
		}
	})
}

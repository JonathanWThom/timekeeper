package main

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
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

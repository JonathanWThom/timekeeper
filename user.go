package main

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// User lives on the users table
// token does not live in the db, but is returned with the user object when needed
type User struct {
	ID       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Token    string `json:"token"`
}

func (user *User) getToken() (string, error) {
	signer := jwt.New(jwt.GetSigningMethod("RS256"))

	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(24))
	claims["iat"] = time.Now().Unix()
	claims["userID"] = user.ID
	signer.Claims = claims

	token, err := signer.SignedString(signKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

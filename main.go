package main

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

type MyCustomClaims struct {
	Foo string `json:"foo"`
	jwt.StandardClaims
}

func main() {
	mySigningKey := []byte("AllYourBase")
	claims := MyCustomClaims{
		"bar",
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("%v", ss)
}

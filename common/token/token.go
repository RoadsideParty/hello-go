package token

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

func GetToken() string {
	key := []byte{1, 23, 4, 5}
	t := jwt.New(jwt.SigningMethodHS256)
	s, err := t.SignedString(key)
	if err != nil {
		fmt.Print(err)
	}
	return s
}

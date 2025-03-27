package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var mySecretKey = []byte("secret")

func ParseToken(tokenString string) (*jwt.Token, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("unexpected signing method")
        }
        return mySecretKey, nil
    })
    return token, err
}

func GenerateToken() (string, error) {
    claims := jwt.MapClaims{
        "username": "user1",
        "exp":      jwt.TimeFunc().Add(1 * time.Hour).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(mySecretKey)
}

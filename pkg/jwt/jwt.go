package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(id int64, username, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":       id,
			"username": username,
			"exp":      time.Now().Add(10 * time.Minute).Unix(),
		},
	)

	key := []byte(secretKey)
	tokenSrt, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenSrt, nil
}

func ValidateToken(tokenStr, secretKey string) (int64, string, error) {
	key := []byte(secretKey)
	clains := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, clains, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return 0, "", err
	}
	if !token.Valid {
		return 0, "", errors.New("invalid Token")
	}
	return int64(clains["id"].(float64)), clains["username"].(string), nil
}

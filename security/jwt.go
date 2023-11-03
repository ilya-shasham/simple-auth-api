package security

import (
	"auth-api/globals"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJwt(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS512)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().UnixNano() + globals.JwtTokenExpiry
	claims["username"] = username
	result, err := token.SignedString(globals.JwtSecret)
	return result, err
}

func CheckJwt(raw string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(raw, func(t *jwt.Token) (interface{}, error) {
		return globals.JwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims := token.Claims.(jwt.MapClaims)

	if claims["exp"].(float64) < float64(time.Now().UnixNano()) {
		return nil, errors.New("jwt token expired")
	}

	username, ok := claims["username"].(string)

	if !ok || username == "" {
		return nil, errors.New("bad username")
	}

	return claims, nil
}

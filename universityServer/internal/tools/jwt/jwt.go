package jwtActions

import (
	"time"
	db "universityServer/internal/database"

	"github.com/golang-jwt/jwt"
)

func CreateJWT(username string, id int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["name"] = username
	claims["id"] = id
	claims["time"] = time.Now().Unix()

	key, err := getKey()

	if err != nil {
		return "", err
	}

	tokenString, err := token.SignedString(key)

	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func getKey() (string, error) {
	key, err := db.GetKey(db.ConnectDB())

	if err != nil {
		return "", err
	}
	return key, nil
}

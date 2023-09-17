package jwtActions

import (
	"errors"
	"fmt"
	"net/http"
	"time"
	db "universityServer/internal/database"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type jwtClaims struct {
	jwt.StandardClaims
	string
	int
	int64
}

func CreateJWT(username string, id int) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * 10).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    username,
		},
		username,
		id,
		time.Now().Unix(),
	})

	key, err := getKey()

	if err != nil {
		return "", err
	}

	tokenString, err := token.SignedString([]byte(key))

	if err != nil {
		return "", err
	}
	fmt.Println(tokenString)
	return tokenString, nil

}

func getKey() (string, error) {
	key, err := db.GetKey(db.ConnectDB())

	if err != nil {
		return "", err
	}
	return key, nil
}

func ValidationJWT(innerFunc func(ctx echo.Context)) func(ctx echo.Context) {
	return func(c echo.Context) {
		if c.Request().Header["Token"] != nil {
			token, err := jwt.Parse(c.Request().Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					c.String(http.StatusUnauthorized, "not authorized")
					return nil, errors.New("Not authorized")
				}
				key, err := getKey()
				if err != nil {
					return nil, err
				}

				return key, nil

			})

			if err != nil {
				c.String(http.StatusUnauthorized, "not authorized")
			}

			if token.Valid {
				innerFunc(c)
			}
		} else {
			c.String(http.StatusUnauthorized, "not authorized")
		}
	}
}

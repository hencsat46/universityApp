package jwtActions

import (
	"errors"
	"net/http"
	"time"
	db "universityServer/internal/database"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type jwtClaims struct {
	jwt.StandardClaims
	username string
	id       string
	int64
}

func CreateJWT(username string, id string, expTime int) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(expTime)).Unix(),
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
	return tokenString, nil

}

func getKey() (string, error) {
	key, err := db.GetKey(db.ConnectDB())

	if err != nil {
		return "", err
	}
	return key, nil
}

func ValidationJWT(innerFunc func(ctx echo.Context) error, giveToken func(ctx echo.Context) error) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		if c.Request().Header["Token"] != nil && c.Request().Header["Token"][0] != "null" {
			token, err := jwt.Parse(c.Request().Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					c.String(http.StatusUnauthorized, "Токен говно")
					return nil, errors.New("not authorized")
				}
				key, err := getKey()
				if err != nil {
					c.String(200, "aaaaaa")
					return nil, err
				}

				return []byte(key), nil

			})

			if err != nil {
				c.String(200, "Неверный токен")
				return err
			}

			if token.Valid {
				innerFunc(c)
				return nil
			}
			return nil
		} else {
			giveToken(c)
			return nil
		}
	})
}

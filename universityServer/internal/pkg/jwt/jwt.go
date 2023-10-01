package jwtActions

import (
	"errors"
	"fmt"
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
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(expTime)).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    username,
		},
		username,
		id,
		time.Now().Unix(),
	})

	key, err := GetKey()

	if err != nil {
		return "", err
	}

	tokenString, err := token.SignedString([]byte(key))

	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func GetKey() (string, error) {
	key, err := db.GetKey(db.ConnectDB())

	if err != nil {
		return "", err
	}
	return key, nil
}

func ValidationJWT(innerFunc func(ctx echo.Context) error) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		if c.Request().Header["Token"] != nil && c.Request().Header["Token"][0] != "null" {
			token, err := jwt.Parse(c.Request().Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					fmt.Println("Токен говно")
					return nil, errors.New("not authorized")
				}
				key, err := GetKey()
				if err != nil {
					fmt.Println(err)
					return nil, err
				}

				return []byte(key), nil

			})

			if err != nil {
				fmt.Println(err, "хуйня")
				return err
			}

			if token.Valid {
				innerFunc(c)
				return nil
			}
			return nil
		} else {
			fmt.Println("no tokens")
			return nil
		}
	})
}

func GetUsernameFromToken(token string) (string, error) {
	claims := jwt.MapClaims{}
	secretKey, err := GetKey()
	if err != nil {
		return "", err
	}
	jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {

		return []byte(secretKey), nil
	})

	username := fmt.Sprint(claims["iss"])

	return username, nil

}

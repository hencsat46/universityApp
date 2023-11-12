package jwtActions

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"universityServer/internal/models"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type jwtClaims struct {
	jwt.StandardClaims
	username string
	id       uint
	int64
}

func CreateJWT(username string, id uint, expTime int) (string, error) {

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

	key := os.Getenv("JWT_KEY")

	tokenString, err := token.SignedString([]byte(key))

	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func ValidationJWT(innerFunc func(ctx echo.Context) error) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		if c.Request().Header["Token"] != nil && c.Request().Header["Token"][0] != "null" {
			token, err := jwt.Parse(c.Request().Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					log.Println("Токен говно")
					return nil, errors.New("not authorized")
				}
				key := os.Getenv("JWT_KEY")

				return []byte(key), nil

			})

			if err != nil {
				log.Println(err)
				return c.JSON(http.StatusUnauthorized, models.Response{Status: http.StatusUnauthorized, Payload: "Authentification error"})
			}

			if token.Valid {
				log.Println("Token valid")
				innerFunc(c)
				// return c.JSON(http.StatusOK, models.Response{Status: http.StatusOK, Payload: "Sign in ok"})
			}
			return nil
		} else {
			log.Println("no token in header")
			return c.JSON(http.StatusUnauthorized, models.Response{Status: http.StatusUnauthorized, Payload: "Authentification error"})
		}
	})
}

func GetUsernameFromToken(token string) (string, error) {
	claims := jwt.MapClaims{}
	secretKey := os.Getenv("JWT_KEY")

	jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {

		return []byte(secretKey), nil
	})

	username := fmt.Sprint(claims["iss"])

	return username, nil

}

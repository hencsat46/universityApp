package jwtActions

import (
	"net/http"
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

func ValidationJWT(innerFunc func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(write http.ResponseWriter, read *http.Request) {
		if read.Header["Token"] != nil {
			token, err := jwt.Parse(read.Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					write.WriteHeader(http.StatusUnauthorized)
					write.Write([]byte("not authorized"))
				}
				key, err := getKey()
				if err != nil {
					return nil, err
				}

				return key, nil

			})

			if err != nil {
				write.WriteHeader(http.StatusUnauthorized)
				write.Write([]byte("not authorized"))
			}

			if token.Valid {
				innerFunc(write, read)
			}
		} else {
			write.WriteHeader(http.StatusUnauthorized)
			write.Write([]byte("not authorized"))
		}
	})
}

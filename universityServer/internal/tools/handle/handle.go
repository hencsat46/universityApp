package handle

import (
	"errors"
	"fmt"
	"strconv"
	database "universityServer/internal/database"
	jwtToken "universityServer/internal/pkg/jwt"
)

func ParseUniversityJson(number int) ([]string, error) {
	result, err := database.GetUniversity(number)
	if err != nil {
		fmt.Println(err)
		return result, nil
	}

	return result, nil

}

func SignUp(user map[string]string) (string, error) {
	var username, password string
	if user["username"] == "" || user["password"] == "" {
		return "", errors.New("Incorrect login or password")
	}

	username = user["username"]
	password = user["password"]
	stringId, err := database.SetUser(username, password)

	if err != nil {
		return "", err
	}

	id, err := strconv.Atoi(stringId)
	token, err := jwtToken.CreateJWT(username, id)

	if err != nil {
		return "", err
	}

	return token, nil

}

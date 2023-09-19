package handle

import (
	"errors"
	"fmt"
	"strings"
	database "universityServer/internal/database"
	jwtToken "universityServer/internal/pkg/jwt"

	"github.com/labstack/echo/v4"
)

func ParseUniversityJson(number int) ([]string, error) {
	result, err := database.GetUniversity(number)
	if err != nil {
		fmt.Println(err)
		return result, nil
	}

	return result, nil

}

func checkEmpty(userData map[string]string) (bool, string, string) {
	if userData["username"] == "" || userData["password"] == "" {
		return false, "", ""
	}

	return true, userData["username"], userData["password"]
}

func SignIn(user map[string]string, expTime int) (string, error) {

	check, username, password := checkEmpty(user)

	if !check {
		return "", errors.New("username or password is empty")
	}

	err := database.Authorization(username, password)

	if err != nil {
		return "", nil
	}

	userId, err := database.GetId(username)
	if err != nil {
		return "", nil
	}

	newToken, err := jwtToken.CreateJWT(username, userId, expTime)
	if err != nil {
		return "", err
	}

	return newToken, nil

}

func SignUp(user map[string]string) error {

	check, username, password := checkEmpty(user)

	if !check {
		return errors.New("username or password is empty")
	}

	err := database.SetUser(username, password)

	if err != nil {
		return err
	}

	return nil

}

func ErrorHandler(ctx echo.Context, err error) int {
	errString := strings.Split(err.Error(), "'")[0]
	switch errString {
	case "invalid character":
		return 0
	case "signature is invalid":
		return 1
	case "Token is expired":
		return 2
	}
	return 0
}

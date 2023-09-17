package handle

import (
	"errors"
	"fmt"
	database "universityServer/internal/database"
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

func SignIn(user map[string]string) error {

	check, username, password := checkEmpty(user)

	if !check {
		return errors.New("Username or password is empty")
	}

	err := database.Authorization(username, password)

	if err != nil {
		return nil
	}

	return nil

}

func SignUp(user map[string]string) error {

	check, username, password := checkEmpty(user)

	if !check {
		return errors.New("Username or password is empty")
	}

	err := database.SetUser(username, password)

	if err != nil {
		return err
	}

	return nil

}

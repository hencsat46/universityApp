package handle

import (
	"errors"
	"fmt"
	database "universityServer/internal/database"
	jwtToken "universityServer/internal/pkg/jwt"
)

func ParseUniversityJson(number int) (map[string]string, error) {
	result, err := database.GetUniversity(number)
	if err != nil {
		fmt.Println(err)
		return result, nil
	}

	return result, nil

}

func GetRemain() (string, error) {
	remain, err := database.GetRemain()
	if err != nil {
		return "", err
	}

	return remain, nil
}

func checkEmpty(username string, password string) bool {
	if username == "" || password == "" {
		return false
	}

	return true
}

func SignIn(user map[string]string, expTime int) (string, error) {

	username, password := user["Username"], user["Password"]

	check := checkEmpty(username, password)

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

	if username == "admin" {
		fmt.Println("aaaaaaaaaaaaa")
		return "", nil
	}

	newToken, err := jwtToken.CreateJWT(username, userId, expTime)
	if err != nil {
		return "", err
	}

	return newToken, nil

}

func ParseStudentRequest(data map[string]string) error {
	username, studentUniversity, points := data["Username"], data["University"], data["Points"]

	err := database.AddStudentRecord(username, studentUniversity, points)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func SignUp(user map[string]string) error {

	studentName, studentSurname, username, password := user["StudentName"], user["StudentSurname"], user["Username"], user["Password"]
	fmt.Println(username, password)
	check := checkEmpty(username, password)

	if !check {
		return errors.New("username or password is empty")
	}

	err := database.SetUser(username, password, studentName, studentSurname)

	if err != nil {
		return err
	}

	return nil

}

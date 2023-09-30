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

func ParseRecords() (map[string]string, error) {

	records, err := database.GetRecords()

	if err != nil {
		fmt.Println(err)
		return make(map[string]string), err
	}

	length := len(records)

	recordsMap := make(map[string]string)

	var finalString string

	for i := 0; i < length; i++ {
		tempString := fmt.Sprintf("%s|%s|%s|%s", records[i][0], records[i][1], records[i][2], records[i][3])
		finalString += tempString
		if length-1 != i {
			finalString += ";"
		}

	}

	recordsMap["records"] = finalString

	return recordsMap, nil

}

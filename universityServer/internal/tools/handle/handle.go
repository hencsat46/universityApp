package handle

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	database "universityServer/internal/database"
	"universityServer/internal/models"
	jwtToken "universityServer/internal/pkg/jwt"
)

func ParseUniversityJson(number int) []models.Universities {
	result := database.ReadUniversity(number)

	log.Println(result)
	return result

}

func GetRemain() (int64, error) {
	remain, err := database.GetRemain()

	return remain, err
}

func checkEmpty(username string, password string) bool {
	if username == "" || password == "" {
		return false
	}

	return true
}

func Ping() (bool, error) {
	if status, err := database.GetStatus(); err != nil {
		log.Println(err)
		return false, err
	} else {
		log.Println(status)
		return status, nil
	}
}

func SignIn(username string, password string, expTime int) (string, error) {

	check := checkEmpty(username, password)

	if !check {
		return "", errors.New("username or password is empty")
	}

	err := database.SignIn(username, password)

	if err != nil {
		log.Println(err)
		return "", err
	}

	userId, err := database.GetId(username)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(username)
	if username == "admin" {
		log.Println("Is admin")
		return "", nil
	}

	newToken, err := jwtToken.CreateJWT(username, userId, expTime)
	if err != nil {
		return "", err
	}

	return newToken, nil

}

func ParseStudentRequest(username, studentUniversity, points string) error {
	status, err := strconv.ParseBool(os.Getenv("DOC_STATUS"))
	if err != nil {
		log.Println(err)
		return err
	}
	if status {
		if err = database.AddStudentRecord(username, studentUniversity, points); err != nil {
			return err
		} else {
			return nil
		}
	} else {
		log.Println(status)
		return errors.New("submission of documents ended")
	}

}

func SignUp(studentName, studentSurname, username, password string) error {

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

func ParseRecords() ([]models.StudentInfo, error) {

	arr, err := database.GetRecords()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return arr, nil

	// length := len(records)

	// recordsMap := make(map[string]string)

	// var finalString string

	// for i := 0; i < length; i++ {
	// 	tempString := fmt.Sprintf("%s|%s|%s|%s", records[i][0], records[i][1], records[i][2], records[i][3])
	// 	finalString += tempString
	// 	if length-1 != i {
	// 		finalString += ";"
	// 	}

	// }

	// recordsMap["records"] = finalString

	// return recordsMap, nil

}

func EditSend(data, user string) error {
	if user != "admin" {
		return errors.New("permission denied")
	}
	var status bool
	switch data {
	case "Продолжить":
		status = true
	case "Остановить":
		status = false
	default:
		return errors.New("wrong json format")
	}

	if err := os.Setenv("DOC_STATUS", strconv.FormatBool(status)); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func GetStudentInfo(username string) (map[string]string, error) {
	studentData, err := database.GetInfoDb(username)

	if err != nil {
		fmt.Println(err)
		return make(map[string]string), err
	}

	studentMap := make(map[string]string)
	studentMap["Username"] = studentData[0]
	studentMap["Name"] = studentData[1]
	studentMap["Surname"] = studentData[2]
	studentMap["University"] = studentData[3]

	return studentMap, nil
}

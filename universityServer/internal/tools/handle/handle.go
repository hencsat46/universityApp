package handle

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"universityServer/internal/api/handlers"
	"universityServer/internal/models"
	jwtToken "universityServer/internal/pkg/jwt"
)

type usecase struct {
	repoInterfaces RepostitoryInterfaces
}

type RepostitoryInterfaces interface {
	ReadUniversity() []models.Universities
	GetRemain() (int64, error)
	GetRecords() ([]models.StudentInfo, error)
	AddStudentRecord(string, string, string) error
	GetStatus() (bool, error)
	GetId(string) (uint, error)
	SetUser(string, string, string, string) error
	SignIn(string, string) error
	GetInfoDb(string) (models.StudentInfo, error)
	GetResult() ([]models.ResultRecord, int, error)
	GetUniversityId(string) (int, error)
}

func NewUsecase(repo RepostitoryInterfaces) handlers.UsecaseInterfaces {
	return &usecase{repoInterfaces: repo}
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func (u *usecase) GetRemain() (int64, error) {
	remain, err := u.repoInterfaces.GetRemain()
	if err != nil {
		log.Println(err)
		return -1, err
	}

	return remain, nil
}

func (u *usecase) ParseUniversityJson() []models.Universities {
	result := u.repoInterfaces.ReadUniversity()

	return result

}

func checkEmpty(username string, password string) bool {
	if username == "" || password == "" {
		return false
	}

	return true
}

func (u *usecase) Ping(username string) (int, error) {
	universityId, err := u.repoInterfaces.GetUniversityId(username)
	return universityId, err
}

func (u *usecase) SignIn(username string, password string, expTime int) (string, error) {

	check := checkEmpty(username, password)

	if !check {
		return "", errors.New("username or password is empty")
	}

	err := u.repoInterfaces.SignIn(username, password)

	if err != nil {
		log.Println(err)
		return "", err
	}

	userId, err := u.repoInterfaces.GetId(username)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(username)

	newToken, err := jwtToken.CreateJWT(username, userId, expTime)
	if err != nil {
		return "", err
	}

	return newToken, nil

}

func (u *usecase) ParseStudentRequest(username, studentUniversity, points string) error {
	status, err := strconv.ParseBool(os.Getenv("DOC_STATUS"))
	if err != nil {
		log.Println(err)
		return err
	}
	if status {
		if err = u.repoInterfaces.AddStudentRecord(username, studentUniversity, points); err != nil {
			return err
		} else {
			return nil
		}
	} else {
		log.Println(status)
		return errors.New("submission of documents ended")
	}

}

func (u *usecase) SignUp(studentName, studentSurname, username, password string) error {

	check := checkEmpty(username, password)

	if !check {
		return errors.New("username or password is empty")
	}

	err := u.repoInterfaces.SetUser(username, password, studentName, studentSurname)

	if err != nil {
		return err
	}

	return nil

}

func (u *usecase) ParseRecords() ([]models.StudentInfo, error) {

	arr, err := u.repoInterfaces.GetRecords()

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

func (u *usecase) EditSend(data, user string) error {
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

	log.Println("Before editing status", status)

	if err := os.Setenv("DOC_STATUS", strconv.FormatBool(status)); err != nil {
		log.Println(err)
		return err
	}

	log.Println("New env name", os.Getenv("DOC_STATUS"))

	return nil
}

func (u *usecase) GetStudentInfo(tokenHeader string) (models.StudentInfo, error) {
	username, err := jwtToken.GetUsernameFromToken(tokenHeader)

	log.Println("Username of profile requester", username)

	if err != nil {
		log.Println(err)
		return models.StudentInfo{}, err
	}
	studentData, err := u.repoInterfaces.GetInfoDb(username)

	if err != nil {
		log.Println(err)
		return models.StudentInfo{}, err
	}

	return studentData, nil
}

func (u *usecase) GetResult() ([]models.ResultRecord, error) {

	status, err := strconv.ParseBool(os.Getenv("DOC_STATUS"))

	if err != nil {
		return nil, err
	}

	if status {
		return nil, nil
	}

	result, emptyCount, err := u.repoInterfaces.GetResult()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	finalResult := make([]models.ResultRecord, 0, len(result)-emptyCount)

	for _, value := range result {
		if len(value.Student_information.([]models.ResultStudent)) != 0 {
			//log.Println("Empty information")
			finalResult = append(finalResult, value)
		}
	}

	return finalResult, nil

}

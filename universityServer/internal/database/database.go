package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	db "universityServer/internal/migrations"
	"universityServer/internal/models"

	pgx "github.com/jackc/pgx/v5"
	"gorm.io/gorm"
)

func ConnectDB() *pgx.Conn {
	config, _ := pgx.ParseConfig(os.Getenv("DB_URL"))

	conn, err := pgx.ConnectConfig(context.Background(), config)

	if err != nil {
		log.Fatal("Cannot to connect to database\n", err)
		return nil
	}

	return conn
}

func ReadUniversity(border int) []models.Universities {
	uni := make([]models.Universities, 2)

	log.Println(db.DB)
	if err := db.DB.Offset(border).Limit(2).Select("Uni_name", "Uni_des", "Uni_img").Find(&uni).Error; err != nil {
		log.Println(err)
	}

	return uni

}

func GetRemain() (int64, error) {

	var remain int64

	if err := db.DB.Model(&models.Universities{}).Count(&remain).Error; err != nil {
		log.Println(err)
		return -1, err
	}

	return remain, nil

}

func GetRecords() ([][]string, error) {
	conn := ConnectDB()
	defer conn.Close(context.Background())

	countQuery, err := conn.Query(context.Background(), "SELECT COUNT(*) FROM students_records;")

	if err != nil {
		fmt.Println(err)
		return make([][]string, 0), err
	}

	var count int
	countQuery.Next()
	countQuery.Scan(&count)
	countQuery.Close()

	recordsArr := make([][]string, 0, count)

	recordsQuery, err := conn.Query(context.Background(), "SELECT * FROM get_records();")

	if err != nil {
		fmt.Println(err)
		return make([][]string, 0), err
	}

	for recordsQuery.Next() {
		record := make([]string, 4)
		recordsQuery.Scan(&record[0], &record[1], &record[2], &record[3])
		recordsArr = append(recordsArr, record)
	}

	return recordsArr, nil

}

func AddStudentRecord(studentName string, studentUniversity string, studentPoints string) error {
	conn := ConnectDB()
	defer conn.Close(context.Background())

	result, err := conn.Query(context.Background(), fmt.Sprintf("SELECT * FROM add_record('%s', '%s', %s);", studentName, studentUniversity, studentPoints))
	if err != nil {
		return err
	}
	var status string
	result.Next()
	err = result.Scan(&status)
	if err != nil {
		return err
	}
	if status == "0" {
		return nil
	}
	return errors.New("student or university doesn't exists")
}

func ChangeStatus(status string) error {
	conn := ConnectDB()
	defer conn.Close(context.Background())

	_, err := conn.Query(context.Background(), fmt.Sprintf("UPDATE records_status SET status = '%s';", status))

	if err != nil {
		return err
	}

	return nil
}

func GetStatus() (bool, error) {
	conn := ConnectDB()
	defer conn.Close(context.Background())

	queryStatus, err := conn.Query(context.Background(), "SELECT status FROM records_status;")
	queryStatus.Next()
	if err != nil {
		return false, err
	}
	var stringStatus string
	queryStatus.Scan(stringStatus)
	status, _ := strconv.ParseBool(stringStatus)
	defer queryStatus.Close()
	return status, nil

}

func GetId(username string) (string, error) {
	conn := ConnectDB()
	defer conn.Close(context.Background())

	var userId string

	response, err := conn.Query(context.Background(), fmt.Sprintf("SELECT user_id FROM users WHERE username = '%s';", username))
	if err != nil {
		return "", err
	}
	response.Next()
	err = response.Scan(&userId)
	if err != nil {
		return "", err
	}
	return userId, nil
}

func SetUser(username string, password string, studentName string, studentSurname string) error {

	user := models.Users{Username: username, Passwd: password, Student_name: studentName, Student_surname: studentSurname}
	log.Println(user)

	if status, err := checkUser(db.DB, &user); err != nil {
		return err
	} else {
		if status == 1 {
			return errors.New("this username already exists")
		}
	}

	if err := db.DB.Omit("User_id").Create(&user).Error; err != nil {
		log.Println(err)
		return err
	}

	// conn := ConnectDB()
	// defer conn.Close(context.Background())
	// var count string
	// response, err := conn.Query(context.Background(), fmt.Sprintf("SELECT * FROM checkUser('%s');", username))

	// if err != nil {
	// 	return err
	// }

	// response.Next()
	// response.Scan(&count)
	// if count != "-1" {
	// 	return errors.New("this user already exists")
	// }
	// response.Close()

	// response, err = conn.Query(context.Background(), fmt.Sprintf("INSERT INTO users(username, passwd, student_name, student_surname) VALUES ('%s', '%s', '%s', '%s')", username, password, studentName, studentSurname))

	// if err != nil {
	// 	return err
	// }
	// defer response.Close()

	return nil
}

func checkUser(database *gorm.DB, model *models.Users) (int64, error) {

	var result int64

	if err := database.Model(&model).Where("username = ?", model.Username).Count(&result).Error; err != nil {
		return -1, err
	}

	return result, nil

}

func Authorization(username string, password string) error {
	conn := ConnectDB()
	defer conn.Close(context.Background())

	response, err := conn.Query(context.Background(), fmt.Sprintf("SELECT * FROM login('%s', '%s');", username, password))

	if err != nil {
		fmt.Println(err)
		return err
	}

	response.Next()
	var status string

	response.Scan(&status)
	fmt.Println(status)

	if status == "0" {
		return nil
	}

	return errors.New("wrong login or password")

}

func GetInfoDb(username string) ([]string, error) {
	conn := ConnectDB()
	defer conn.Close(context.Background())

	studentResponse, err := conn.Query(context.Background(), fmt.Sprintf("SELECT * FROM get_user_data('%s')", username))

	if err != nil {
		fmt.Println(err)
		return make([]string, 0), err
	}

	studentData := make([]string, 4)

	studentResponse.Next()
	studentResponse.Scan(&studentData[0], &studentData[1], &studentData[2], &studentData[3])

	fmt.Println(studentData)

	return studentData, nil

}

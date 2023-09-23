package database

import (
	"context"
	"errors"
	"fmt"
	"log"

	pgx "github.com/jackc/pgx/v5"
)

const DB_URL = "postgresql://postgres:forstudy@localhost:5432/universityDB"

func ConnectDB() *pgx.Conn {
	config, _ := pgx.ParseConfig(DB_URL)

	conn, err := pgx.ConnectConfig(context.Background(), config)

	if err != nil {
		log.Fatal("Cannot to connect to database", err)
		return nil
	}

	return conn
}

func GetUniversity(border int) ([]string, error) {
	conn := ConnectDB()
	defer conn.Close(context.Background())

	result, err := conn.Query(context.Background(), fmt.Sprintf("SELECT uni_name, uni_des, uni_img FROM tempUni OFFSET %v LIMIT %v;", border, border+1))
	result.Next()
	defer func() {
		if err == nil {
			return
		}
		result.Close()
	}()
	universityArray := make([]string, 4)
	if err != nil {
		return universityArray, err
	}

	err = result.Scan(&universityArray[0], &universityArray[1], &universityArray[2])

	if err != nil {
		return universityArray, err
	}

	result.Close()
	universityArray[3], err = GetRemain()
	if err != nil {
		return universityArray, err
	}

	if err != nil {
		return universityArray, err
	}

	return universityArray, nil

}

func GetRemain() (string, error) {
	conn := ConnectDB()
	defer conn.Close(context.Background())
	remained, err := conn.Query(context.Background(), "SELECT COUNT(*) FROM tempUni;")
	defer func() {
		if err == nil {
			return
		}
		remained.Close()
	}()

	if err != nil {
		return "", err
	}
	var remainedString string
	remained.Next()
	err = remained.Scan(&remainedString)

	if err != nil {
		return "", err
	}

	return remainedString, nil

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
	conn := ConnectDB()
	defer conn.Close(context.Background())
	var count string
	response, err := conn.Query(context.Background(), fmt.Sprintf("SELECT * FROM checkUser('%s');", username))

	if err != nil {
		return err
	}

	response.Next()
	response.Scan(&count)
	if count != "-1" {
		return errors.New("this user already exists")
	}
	response.Close()

	response, err = conn.Query(context.Background(), fmt.Sprintf("INSERT INTO users(username, passwd, student_name, student_surname) VALUES ('%s', '%s', '%s', '%s')", username, password, studentName, studentSurname))

	if err != nil {
		return err
	}
	defer response.Close()

	return nil
}

func Authorization(username string, password string) error {
	conn := ConnectDB()
	defer conn.Close(context.Background())

	response, err := conn.Query(context.Background(), fmt.Sprintf("SELECT * FROM login('%s', '%s');", username, password))

	if err != nil {
		return err
	}

	response.Next()
	var nameData, passData string

	response.Scan(&nameData, &passData)

	fmt.Println(nameData, passData)

	return nil

}

func GetKey(conn *pgx.Conn) (string, error) {
	key, err := conn.Query(context.Background(), "SELECT secretKey FROM jwtKey;")

	if err != nil {
		return "", err
	}
	defer key.Close()

	key.Next()
	var stringKey string
	err = key.Scan(&stringKey)
	if err != nil {
		return "", err
	}

	return stringKey, nil
}

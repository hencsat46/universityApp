package database

import (
	"context"
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

	universityArray := make([]string, 4)
	if err != nil {
		return universityArray, err
	}

	err = result.Scan(&universityArray[0], &universityArray[1], &universityArray[2])

	if err != nil {
		return universityArray, err
	}

	result.Next()
	universityArray[3], err = getRemain(conn)
	if err != nil {
		return universityArray, err
	}

	if err != nil {
		return universityArray, err
	}

	return universityArray, nil

}

func SetUser(username string, password string) (string, error) {
	conn := ConnectDB()
	defer conn.Close(context.Background())
	response, err := conn.Query(context.Background(), fmt.Sprintf("INSERT INTO users(username, passwd) VALUES ('%s', '%s')", username, password))
	if err != nil {
		return "-1", err
	}
	response.Next()

	response, err = conn.Query(context.Background(), fmt.Sprintf("SELECT user_id FROM users WHERE username='%s'", username))

	if err != nil {
		return "-1", err
	}

	var id string
	response.Next()
	err = response.Scan(&id)

	if err != nil {
		return "-1", err
	}
	return id, nil
}

func getRemain(conn *pgx.Conn) (string, error) {
	remained, err := conn.Query(context.Background(), "SELECT COUNT(*) FROM tempUni;")

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

func GetKey(conn *pgx.Conn) (string, error) {
	key, err := conn.Query(context.Background(), "SELECT secretKey FROM jwtKey;")

	if err != nil {
		return "", err
	}

	key.Next()
	var stringKey string
	err = key.Scan(&stringKey)
	if err != nil {
		return "", err
	}

	return stringKey, nil
}

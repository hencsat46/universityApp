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

func GetUniversity(border int) {
	conn := ConnectDB()
	defer conn.Close(context.Background())

	result, _ := conn.Query(context.Background(), "SELECT uni_name, uni_des, uni_img FROM tempuni;")

	var p1, p2, p3 string

	result.Scan(&p1, &p2, &p3)

	fmt.Println(p1, p2, p3)

}

package database

import (
	"context"
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

type UniversityUnit struct {
	name        string
	description string
	imgPath     string
}

func GetUniversity() (UniversityUnit, error) {
	conn := ConnectDB()
	defer conn.Close(context.Background())
	var unit UniversityUnit
	result, err := conn.Query(context.Background(), "SELECT * FROM tempUni LIMIT 1")

	if err == nil {
		result.Scan(&unit.name, &unit.description, &unit.imgPath)
		return unit, nil
	}

	return unit, err
}

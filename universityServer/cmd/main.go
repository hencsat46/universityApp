package main

import (
	app "universityServer/internal/api/server"
	db "universityServer/internal/migrations"
	"universityServer/internal/pkg/env"
)

func main() {
	env.Env()
	db.InitDB()

	app.Run()

}

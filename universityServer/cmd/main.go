package main

import (
	"universityServer/internal/api/handlers"
	"universityServer/internal/database"
	db "universityServer/internal/migrations"
	"universityServer/internal/pkg/env"
	"universityServer/internal/tools/handle"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	env.Env()
	db.InitDB()

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "DELETE", "POST", "PUT"},
	}))

	repo := database.NewRepostitory()
	usecase := handle.NewUsecase(repo)
	handler := handlers.NewHandler(usecase)
	handler.Routes(e)

	e.Start("0.0.0.0:3000")

}

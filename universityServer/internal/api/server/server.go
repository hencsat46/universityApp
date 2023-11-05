package server

import (
	router "universityServer/internal/api/routes"
	env "universityServer/internal/pkg/env"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run() {
	e := echo.New()

	env.Env()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "DELETE", "POST", "PUT"},
	}))

	router.Routes(e)
	e.Start("0.0.0.0:3000")
}

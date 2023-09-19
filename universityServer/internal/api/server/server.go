package server

import (
	router "universityServer/internal/api/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "DELETE", "POST", "PUT"},
	}))

	router.Routes(e)
	e.Start(":3000")
}

// func errorIndicator(ctx echo.Context, err error) error {
// 	errorsHandler.ErrorHandler(ctx, err)
// 	return nil
// }

// func expiredToken(ctx echo.Context, err error) error {
// 	if err.Error() == "Token is expired" {

// 	}
// }

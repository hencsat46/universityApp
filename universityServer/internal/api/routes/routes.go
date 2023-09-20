package routes

import (
	handlers "universityServer/internal/api/handlers"
	jwt "universityServer/internal/pkg/jwt"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	e.POST("/getUniversity", handlers.GetUniversity)
	e.POST("/signup", handlers.SignUp)
	e.POST("/signin", handlers.SignIn)
	e.POST("/token", jwt.ValidationJWT(handlers.TokenOk))
	//e.GET("/check", createJwt)

}

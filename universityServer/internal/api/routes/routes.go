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
	e.GET("/token", jwt.ValidationJWT(func(ctx echo.Context) error { return nil }))
	e.GET("/getRemain", handlers.GetRemain)
	e.POST("/addStudent", jwt.ValidationJWT(handlers.AddStudent))
	e.GET("/records", handlers.GetRecords)
	e.POST("/stopSend", jwt.ValidationJWT(handlers.EditSend))
	e.GET("/profile", jwt.ValidationJWT(handlers.UserProfile))
	e.GET("/", jwt.ValidationJWT(handlers.AutoLogin))
	e.GET("/ping", handlers.Ping)
	e.GET("/getresult", handlers.GetResult)

}

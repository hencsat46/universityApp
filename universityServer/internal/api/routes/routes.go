package routes

import (
	"context"
	"log"
	handlers "universityServer/internal/api/handlers"
	db "universityServer/internal/database"
	jwt "universityServer/internal/pkg/jwt"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	e.POST("/getUniversity", handlers.GetUniversity)
	e.POST("/signup", handlers.SignUp)
	e.POST("/signin", handlers.SignIn)
	e.POST("/token", jwt.ValidationJWT(handlers.TokenOk))
	e.GET("/getRemain", handlers.GetRemain)
	e.POST("/addStudent", jwt.ValidationJWT(handlers.AddStudent))
	e.GET("/records", handlers.GetRecords)
	e.POST("/stopSend", handlers.EditSend)
	e.GET("/profile", handlers.UserProfile)
	e.GET("/", jwt.ValidationJWT(handlers.AutoLogin))
	e.GET("/ping", func(ctx echo.Context) error {
		conn := db.ConnectDB()
		err := conn.Ping(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		return ctx.String(200, "haha")
	})

}

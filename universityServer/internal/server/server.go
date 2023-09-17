package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	jwt "universityServer/internal/pkg/jwt"
	usecase "universityServer/internal/tools/handle"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func getUniversity(ctx echo.Context) error {

	dataMap := make(map[string]int)

	err := json.NewDecoder(ctx.Request().Body).Decode(&dataMap)
	if err != nil {
		fmt.Println(err)
		return err
	}

	result, err := usecase.ParseUniversityJson(dataMap["order"])

	if err != nil {
		fmt.Println(err)
		return err
	}

	jsonUniversity := make(map[string]string)
	jsonUniversity["name"] = result[0]
	jsonUniversity["description"] = result[1]
	jsonUniversity["imagePath"] = result[2]
	jsonUniversity["left"] = result[3]
	convertUniversity, err := json.Marshal(jsonUniversity)

	if err != nil {
		fmt.Println(err)
		return err
	}

	ctx.Response().Header().Set("Content-Type", "application/json")
	return ctx.String(http.StatusOK, string(convertUniversity))

}

func registration(ctx echo.Context) error {

	dataMap := make(map[string]string)

	err := json.NewDecoder(ctx.Request().Body).Decode(&dataMap)

	if err != nil {
		fmt.Println(err)
		return err
	}

	jwtToken, err := usecase.SignUp(dataMap)

	if err != nil {
		fmt.Println(err)
		return err
	}

	jsonMap := make(map[string]string)

	jsonMap["token"] = jwtToken

	jsonToken, err := json.Marshal(jsonMap)

	if err != nil {
		fmt.Println(err)
		return err
	}

	ctx.Response().Header().Set("Content-Type", "application/json")

	return ctx.String(http.StatusOK, string(jsonToken))

}

func Run() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "DELETE", "POST", "PUT"},
	}))

	router(e)
	e.Start(":3000")
}

func createJwt(ctx echo.Context) error {
	jwtToken, err := jwt.CreateJWT("ilya", 12)

	if err != nil {
		return err
	}

	return ctx.String(200, jwtToken)
}

func router(e *echo.Echo) {
	e.POST("/getUniversity", getUniversity)
	e.POST("/signup", registration)
	e.GET("/check", createJwt)

}

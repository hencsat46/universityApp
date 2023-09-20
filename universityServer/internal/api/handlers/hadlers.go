package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	usecase "universityServer/internal/tools/handle"

	"github.com/labstack/echo/v4"
)

type ResponseJson struct {
	Status  string
	Message string
}

func SignUp(ctx echo.Context) error {

	dataMap := make(map[string]string)

	err := json.NewDecoder(ctx.Request().Body).Decode(&dataMap)
	fmt.Println(dataMap)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = usecase.SignUp(dataMap)

	if err != nil {
		fmt.Println(err)
		return err
	}

	responseJson := ResponseJson{
		"ok",
		"Sign Up success",
	}

	return ctx.JSON(200, responseJson)

}

func SignIn(ctx echo.Context) error {
	dataMap := make(map[string]string)

	err := json.NewDecoder(ctx.Request().Body).Decode(&dataMap)
	fmt.Println(dataMap)
	if err != nil {
		return err
	}

	expTime := 2

	token, err := usecase.SignIn(dataMap, expTime)

	if err != nil {
		return err
	}

	jsonMap := make(map[string]string)
	jsonMap["Token"] = token

	jsonString, err := json.Marshal(jsonMap)
	if err != nil {
		return err
	}

	return ctx.String(http.StatusOK, string(jsonString))
}

func GetUniversity(ctx echo.Context) error {

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

func TokenOk(ctx echo.Context) error {
	fmt.Println("hello")
	return nil
}

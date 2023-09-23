package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"universityServer/internal/database"
	jsonResponse "universityServer/internal/pkg/responseJson"
	usecase "universityServer/internal/tools/handle"

	"github.com/labstack/echo/v4"
)

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

	responseMap := make(map[string]string)
	jsonStruct, err := jsonResponse.Response(responseMap, "sign up")
	if err != nil {
		fmt.Println(err)
		return err
	}

	return ctx.JSON(200, jsonStruct)

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

	responseMap := make(map[string]string)
	responseMap["Token"] = token
	jsonStruct, err := jsonResponse.Response(responseMap, "sign in")

	if err != nil {
		fmt.Println(err)
		return err
	}

	return ctx.JSON(http.StatusOK, jsonStruct)
}

func GetRemain(ctx echo.Context) error {
	remain, err := database.GetRemain()
	if err != nil {
		jsonStruct, err := jsonResponse.Response(make(map[string]string), err.Error())
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		return ctx.JSON(500, jsonStruct)
	}
	responseMap := make(map[string]string)
	responseMap["remain"] = remain
	jsonStruct, err := jsonResponse.Response(responseMap, "remain")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return ctx.JSON(200, jsonStruct)
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

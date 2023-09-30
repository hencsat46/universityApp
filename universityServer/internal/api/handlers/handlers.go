package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	database "universityServer/internal/database"
	jwtActions "universityServer/internal/pkg/jwt"
	jsonResponse "universityServer/internal/pkg/responseJson"
	usecase "universityServer/internal/tools/handle"

	"github.com/golang-jwt/jwt"
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
	fmt.Println("delivery message")
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

	expTime := 30

	token, err := usecase.SignIn(dataMap, expTime)

	if err != nil {
		return err
	}

	responseMap := make(map[string]string)
	responseMap["Token"] = token
	fmt.Println(responseMap)
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

	if err != nil {
		fmt.Println(err)
		return err
	}
	response, err := jsonResponse.Response(result, "university")
	if err != nil {
		fmt.Println(err)
	}
	ctx.Response().Header().Set("Content-Type", "application/json")
	return ctx.JSON(http.StatusOK, response)

}

func AddStudent(ctx echo.Context) error {
	requestMap := make(map[string]string)
	claims := jwt.MapClaims{}
	secretKey, err := jwtActions.GetKey()
	if err != nil {
		return err
	}
	jwt.ParseWithClaims(ctx.Request().Header["Token"][0], claims, func(token *jwt.Token) (interface{}, error) {

		return []byte(secretKey), nil
	})

	username := fmt.Sprint(claims["iss"])

	err = json.NewDecoder(ctx.Request().Body).Decode(&requestMap)

	if err != nil {
		fmt.Println(err)
		return err
	}

	requestMap["Username"] = username

	usecase.ParseStudentRequest(requestMap)

	responseJson, err := jsonResponse.Response(requestMap, "add student")

	if err != nil {
		return err
	}

	return ctx.JSON(200, responseJson)
}

func TokenOk(ctx echo.Context) error {
	fmt.Println("hello")
	return nil
}

func GetRecords(ctx echo.Context) error {
	records, err := usecase.ParseRecords()

	if err != nil {
		fmt.Println(err)
		return err
	}

	jsonData, err := jsonResponse.Response(records, "studentRecords")

	if err != nil {
		fmt.Println(err)
		return err
	}

	return ctx.JSON(200, jsonData)
}

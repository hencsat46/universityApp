package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	database "universityServer/internal/database"
	"universityServer/internal/models"
	jwtActions "universityServer/internal/pkg/jwt"
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
		return ctx.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusInternalServerError, Payload: "database error"})
	}

	return ctx.JSON(http.StatusOK, models.Response{Status: http.StatusOK, Payload: remain})

}

func GetUniversity(ctx echo.Context) error {

	var requestBody GetUniversityDTO = GetUniversityDTO{-1}

	if err := ctx.Bind(&requestBody); err != nil || requestBody.Order == -1 {
		return ctx.JSON(http.StatusBadRequest, models.Response{Status: http.StatusBadRequest, Payload: "wrong json format"})
	}

	result := usecase.ParseUniversityJson(requestBody.Order)

	return ctx.JSON(http.StatusOK, models.Response{Status: http.StatusOK, Payload: result})

}

func AddStudent(ctx echo.Context) error {
	requestMap := make(map[string]string)

	username, err := jwtActions.GetUsernameFromToken(ctx.Request().Header["Token"][0])

	if err != nil {
		fmt.Println(err)
		return err
	}

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

func EditSend(ctx echo.Context) error {
	requestMap := make(map[string]string)

	err := json.NewDecoder(ctx.Request().Body).Decode(&requestMap)

	if err != nil {
		fmt.Println(err)
		return err
	}

	err = usecase.EditSend(requestMap)

	if err != nil {
		fmt.Println(err)
		return err
	}

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

func UserProfile(ctx echo.Context) error {
	token, err := jwtActions.GetUsernameFromToken(ctx.Request().Header["Token"][0])

	if err != nil {
		fmt.Println(err)
		return err
	}

	responseMap, err := usecase.GetStudentInfo(token)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	jsonData, err := jsonResponse.Response(responseMap, "profile")

	if err != nil {
		fmt.Println(err)
		return err
	}

	return ctx.JSON(200, jsonData)
}

func AutoLogin(ctx echo.Context) error {

	jsonData, err := jsonResponse.Response(make(map[string]string), "autoLogin")
	if err != nil {
		fmt.Println(err)
		return err
	}

	return ctx.JSON(200, jsonData)
}

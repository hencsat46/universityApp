package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	database "universityServer/internal/database"
	"universityServer/internal/models"
	jwtActions "universityServer/internal/pkg/jwt"
	usecase "universityServer/internal/tools/handle"

	"github.com/labstack/echo/v4"
)

func SignUp(ctx echo.Context) error {

	var requestBody SignUpDTO = SignUpDTO{"", "", "", ""}

	if err := ctx.Bind(&requestBody); err != nil || requestBody.Username == "" || requestBody.Password == "" {
		return ctx.JSON(http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Payload: "wrong json format"})
	}

	// dataMap := make(map[string]string)

	// err := json.NewDecoder(ctx.Request().Body).Decode(&dataMap)
	// fmt.Println(dataMap)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }

	err := usecase.SignUp(requestBody.StudentName, requestBody.StudentSurname, requestBody.Username, requestBody.Password)
	fmt.Println("delivery message")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Payload: "Username in taken"})
	}

	return ctx.JSON(200, &models.Response{Status: http.StatusOK, Payload: "Sign up ok"})

}

func SignIn(ctx echo.Context) error {

	var requestBody SignInDTO = SignInDTO{"", ""}
	var expTime int

	if t, err := strconv.Atoi(os.Getenv("EXP_TIME")); err != nil {
		return ctx.JSON(http.StatusInternalServerError, &models.Response{Status: http.StatusInternalServerError, Payload: "Internal server error"})
	} else {
		expTime = t
	}

	if err := ctx.Bind(&requestBody); err != nil {
		return ctx.JSON(http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Payload: "Wrong data"})
	}

	if token, err := usecase.SignIn(requestBody.Username, requestBody.Password, expTime); err != nil {
		return ctx.JSON(http.StatusInternalServerError, &models.Response{Status: http.StatusInternalServerError, Payload: "Internal server error"})
	} else {
		return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: token})
	}

	//test on front

}

func GetRemain(ctx echo.Context) error {
	remain, err := database.GetRemain()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &models.Response{Status: http.StatusInternalServerError, Payload: "database error"})
	}

	return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: remain})

}

func Ping(ctx echo.Context) error {
	usecase.Ping()
	return nil
}

func GetUniversity(ctx echo.Context) error {

	var requestBody GetUniversityDTO = GetUniversityDTO{-1}

	if err := ctx.Bind(&requestBody); err != nil || requestBody.Order == -1 {
		return ctx.JSON(http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Payload: "wrong json format"})
	}

	result := usecase.ParseUniversityJson(requestBody.Order)

	return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: result})

}

func AddStudent(ctx echo.Context) error {

	username, err := jwtActions.GetUsernameFromToken(ctx.Request().Header["Token"][0])
	if err != nil {
		fmt.Println(err)
		return err
	}

	requestBody := UniversityStudentDTO{"", ""}

	if err := ctx.Bind(&requestBody); err != nil || requestBody.University == "" || requestBody.Points == "" {
		return ctx.JSON(http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Payload: "wrong json format"})
	}
	if err = usecase.ParseStudentRequest(username, requestBody.University, requestBody.Points); err != nil {

		if err.Error() == "submission of documents ended" {
			return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: err.Error()})
		}

		return ctx.JSON(http.StatusInternalServerError, &models.Response{Status: http.StatusInternalServerError, Payload: "Internal Server Error"})
	}

	return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: "Student added or updated"})
}

func EditSend(ctx echo.Context) error {

	request := EditSendDTO{""}

	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Payload: "Json error"})
	}

	username, err := jwtActions.GetUsernameFromToken(ctx.Request().Header["Token"][0])

	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, &models.Response{Status: http.StatusUnauthorized, Payload: err.Error()})
	}

	if err := usecase.EditSend(request.Status, username); err != nil {
		if err.Error() == "permission denied" {
			return ctx.JSON(http.StatusUnauthorized, &models.Response{Status: http.StatusUnauthorized, Payload: err.Error()})
		}
		return ctx.JSON(http.StatusInternalServerError, &models.Response{Status: http.StatusInternalServerError, Payload: err.Error()})
	}

	return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: "Success"})

}

func GetRecords(ctx echo.Context) error {
	arr, err := usecase.ParseRecords()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &models.Response{Status: http.StatusInternalServerError, Payload: "Internal Server Error"})
	}

	return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: arr})

}

func UserProfile(ctx echo.Context) error {

	response, err := usecase.GetStudentInfo(ctx.Request().Header["Token"][0])

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &models.Response{Status: http.StatusInternalServerError, Payload: "Internal Server Error"})
	}

	return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: response})
}

func AutoLogin(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: "Sign in ok"})
}

func GetResult(ctx echo.Context) error {
	result, err := usecase.GetResult()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &models.Response{Status: http.StatusUnauthorized, Payload: "Internal Server Error"})
	}

	return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: result})
}

package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"universityServer/internal/models"
	jwtActions "universityServer/internal/pkg/jwt"

	"github.com/labstack/echo/v4"
)

type handler struct {
	usecase UsecaseInterfaces
}

type UsecaseInterfaces interface {
	ParseUniversityJson() []models.Universities
	Ping(string) (int, error)
	SignIn(string, string, int) (string, error)
	ParseStudentRequest(string, string, string) error
	SignUp(string, string, string, string) error
	ParseRecords() ([]models.StudentInfo, error)
	EditSend(string, string) error
	GetStudentInfo(string) (models.StudentInfo, error)
	GetResult() ([]models.ResultRecord, error)
	GetRemain() (int64, error)
}

func NewHandler(usecase UsecaseInterfaces) *handler {
	return &handler{usecase: usecase}
}

func (h *handler) Routes(e *echo.Echo) {
	e.GET("/get_universities", h.GetUniversity)
	e.POST("/signup", h.SignUp)
	e.POST("/signin", h.SignIn)
	e.GET("/token", jwtActions.ValidationJWT(func(ctx echo.Context) error { return nil }))
	e.GET("/getRemain", h.GetRemain)
	e.POST("/addStudent", jwtActions.ValidationJWT(h.AddStudent))
	e.GET("/records", h.GetRecords)
	e.POST("/stopSend", jwtActions.ValidationJWT(h.EditSend))
	e.GET("/profile", jwtActions.ValidationJWT(h.UserProfile))
	e.GET("/", jwtActions.ValidationJWT(h.AutoLogin))
	e.GET("/ping", jwtActions.ValidationJWT(h.Ping))
	e.GET("/getresult", h.GetResult)
}

func (h *handler) SignUp(ctx echo.Context) error {

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

	err := h.usecase.SignUp(requestBody.StudentName, requestBody.StudentSurname, requestBody.Username, requestBody.Password)
	fmt.Println("delivery message")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Payload: "Username in taken"})
	}

	return ctx.JSON(200, &models.Response{Status: http.StatusOK, Payload: "Sign up ok"})

}

func (h *handler) SignIn(ctx echo.Context) error {

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

	if token, err := h.usecase.SignIn(requestBody.Username, requestBody.Password, expTime); err != nil {
		if err.Error() == "wrong username or password" {
			return ctx.JSON(http.StatusUnauthorized, &models.Response{Status: http.StatusUnauthorized, Payload: "Authentication error"})
		}
		return ctx.JSON(http.StatusInternalServerError, &models.Response{Status: http.StatusInternalServerError, Payload: "Internal server error"})

	} else {
		return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: token})
	}

	//test on front

}

func (h *handler) GetRemain(ctx echo.Context) error {
	remain, err := h.usecase.GetRemain()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &models.Response{Status: http.StatusInternalServerError, Payload: "database error"})
	}

	return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: remain})

}

func (h *handler) Ping(ctx echo.Context) error {

	username, err := jwtActions.GetUsernameFromToken(ctx.Request().Header["Token"][0])
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusUnauthorized, &models.Response{Status: http.StatusUnauthorized, Payload: "nigga"})
	}

	universityId, err := h.usecase.Ping(username)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &models.Response{Status: http.StatusInternalServerError, Payload: "Internal Server Error"})
	}

	return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: universityId})
}

func (h *handler) GetUniversity(ctx echo.Context) error {

	result := h.usecase.ParseUniversityJson()

	return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: result})

}

func (h *handler) AddStudent(ctx echo.Context) error {

	username, err := jwtActions.GetUsernameFromToken(ctx.Request().Header["Token"][0])
	if err != nil {
		fmt.Println(err)
		return err
	}

	requestBody := UniversityStudentDTO{"", ""}

	if err := ctx.Bind(&requestBody); err != nil || requestBody.University == "" || requestBody.Points == "" {
		return ctx.JSON(http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Payload: "wrong json format"})
	}
	if err = h.usecase.ParseStudentRequest(username, requestBody.University, requestBody.Points); err != nil {

		if err.Error() == "submission of documents ended" {
			return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: err.Error()})
		}

		return ctx.JSON(http.StatusInternalServerError, &models.Response{Status: http.StatusInternalServerError, Payload: "Internal Server Error"})
	}

	universityId, err := h.usecase.Ping(username)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &models.Response{Status: http.StatusInternalServerError, Payload: "Internal Server Error"})
	}

	return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: universityId})
}

func (h *handler) EditSend(ctx echo.Context) error {

	request := EditSendDTO{""}

	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, &models.Response{Status: http.StatusBadRequest, Payload: "Json error"})
	}

	username, err := jwtActions.GetUsernameFromToken(ctx.Request().Header["Token"][0])

	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, &models.Response{Status: http.StatusUnauthorized, Payload: err.Error()})
	}

	if err := h.usecase.EditSend(request.Status, username); err != nil {
		if err.Error() == "permission denied" {
			return ctx.JSON(http.StatusUnauthorized, &models.Response{Status: http.StatusUnauthorized, Payload: err.Error()})
		}
		return ctx.JSON(http.StatusInternalServerError, &models.Response{Status: http.StatusInternalServerError, Payload: err.Error()})
	}

	return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: "Success"})

}

func (h *handler) GetRecords(ctx echo.Context) error {
	arr, err := h.usecase.ParseRecords()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &models.Response{Status: http.StatusInternalServerError, Payload: "Internal Server Error"})
	}

	return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: arr})

}

func (h *handler) UserProfile(ctx echo.Context) error {

	response, err := h.usecase.GetStudentInfo(ctx.Request().Header["Token"][0])
	log.Println(response)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &models.Response{Status: http.StatusInternalServerError, Payload: "Internal Server Error"})
	}

	return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: response})
}

func (h *handler) AutoLogin(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: "Sign in ok"})
}

func (h *handler) GetResult(ctx echo.Context) error {
	result, err := h.usecase.GetResult()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &models.Response{Status: http.StatusUnauthorized, Payload: "Internal Server Error"})
	}

	if err == nil && result == nil {
		log.Println("PENIS")
		log.Println([]models.ResultRecord{models.ResultRecord{Student_university: "No records or document submission is not ended", Student_information: nil}})
		return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: []models.ResultRecord{models.ResultRecord{Student_university: "No records or document submission is not ended", Student_information: nil}}})
	}

	return ctx.JSON(http.StatusOK, &models.Response{Status: http.StatusOK, Payload: result})
}

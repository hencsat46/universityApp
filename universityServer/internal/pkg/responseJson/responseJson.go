package responsejson

import (
	"errors"
	"strings"
)

type SignInMessage struct {
	Message string
	Token   string
}

type ResponseJson struct {
	Status  string
	Payload interface{}
}

type UniversityJson struct {
	FirstUni  string
	SecondUni string
	Remain    RemainMessage
}

type SignUpMessage struct {
	Message string
}

type errorMessages struct {
	Message string
}

type RemainMessage struct {
	Message string
}

type AddStudent struct {
	Message  string
	Username string
}

type Records struct {
	Message [][]string
}

type Profile struct {
	Username   string
	Name       string
	Surname    string
	University string
}

func Response(data map[string]string, status string) (ResponseJson, error) {

	switch status {
	case "sign in":
		value, ok := data["Token"]
		if !ok {
			return ResponseJson{}, errors.New("token not found")
		}
		payload := SignInMessage{"Sign in ok", value}
		responseJson := ResponseJson{"Ok", payload}
		return responseJson, nil
	case "sign up":
		payload := SignUpMessage{"Sign up ok"}
		responsejson := ResponseJson{"Ok", payload}
		return responsejson, nil
	case "remain":
		payload := RemainMessage{data["remain"]}
		return ResponseJson{"Ok", payload}, nil
	case "add student":
		payload := AddStudent{"You were added", data["Username"]}
		return ResponseJson{"Ok", payload}, nil
	case "university":
		remain := RemainMessage{data["2"]}
		payload := UniversityJson{data["0"], data["1"], remain}
		return ResponseJson{"Ok", payload}, nil
	case "studentRecords":
		recordsString := strings.Split(data["records"], ";")
		length := len(recordsString)
		recordsArr := make([][]string, length)
		for i := 0; i < length; i++ {
			recordsArr[i] = strings.Split(recordsString[i], "|")
		}
		records := Records{recordsArr}
		return ResponseJson{"Ok", records}, nil
	case "profile":
		if data["University"] == "" {
			data["University"] = "Вы не подавали документы"
		}
		payload := Profile{data["Username"], data["Name"], data["Surname"], data["University"]}
		return ResponseJson{"Ok", payload}, nil
	default:
		payload := errorMessages{status}
		return ResponseJson{"Ok", payload}, nil
	}

}

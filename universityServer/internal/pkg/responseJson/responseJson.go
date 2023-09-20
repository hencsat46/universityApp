package responsejson

import (
	"errors"
)

type SignInMessage struct {
	Message string
	Token   string
}

type ResponseJson struct {
	Status  string
	Payload interface{}
}

type SignUpMessage struct {
	Message string
}

func Response(data map[string]string, status string) (ResponseJson, error) {

	if status == "sign in" {
		value, ok := data["Token"]
		if !ok {
			return ResponseJson{}, errors.New("token not found")
		}
		payload := SignInMessage{"Sign in ok", value}
		responseJson := ResponseJson{"Ok", payload}
		return responseJson, nil
	}

	if status == "sign up" {
		payload := SignUpMessage{"Sign up ok"}
		responsejson := ResponseJson{"Ok", payload}
		return responsejson, nil
	}

	return ResponseJson{}, errors.New("wrong status")

}

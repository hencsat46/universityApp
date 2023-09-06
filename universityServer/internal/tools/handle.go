package handle

import (
	"fmt"
	database "universityServer/internal/database"
)

func ParseUniversityJson(number int) ([]string, error) {
	result, err := database.GetUniversity(number)
	if err != nil {
		fmt.Println(err)
		return result, nil
	}

	return result, nil

}

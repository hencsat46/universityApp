package handle

import (
	"fmt"
	"strconv"
	database "universityServer/internal/database"
)

func ParseUniversityJson(order string) ([]string, error) {
	number, _ := strconv.Atoi(order)
	result, err := database.GetUniversity(number)
	if err != nil {
		fmt.Println(err)
		return result, nil
	}

	return result, nil

}

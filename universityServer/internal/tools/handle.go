package handle

import (
	"strconv"
	database "universityServer/internal/database"
)

func ParseUniversityJson(order string) (string, error) {
	number, _ := strconv.Atoi(order)
	//result, err := database.GetUniversity(number)
	database.GetUniversity(number)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return result, nil
	// }
	result := "ljksdf"
	return result, nil

}

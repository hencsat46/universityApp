package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	usecase "universityServer/internal/tools"

	"github.com/rs/cors"
)

func auth(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		dataMap := make(map[string]string)

		err := json.NewDecoder(r.Body).Decode(&dataMap)

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(dataMap)

	}
}

type jsonUniversity struct {
	name string `json: name`
	des  string `json: description`
	img  string `json: imagePath`
}

func getUniversity(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		dataMap := make(map[string]string)

		err := json.NewDecoder(r.Body).Decode(&dataMap)

		if err != nil {
			fmt.Println(err)
			return
		}

		result, err := usecase.ParseUniversityJson(dataMap["order"])

		if err != nil {
			fmt.Println(err)
			return
		}

		var university jsonUniversity
		university.name = result[0]
		university.des = result[1]
		university.img = result[2]

		convertUniversity, err := json.Marshal(university)

		if err != nil {
			fmt.Println(err)
			return
		}

		w.Write(convertUniversity)

	}

}

func main() {
	r := http.NewServeMux()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "DELETE", "POST", "PUT"},
	})

	r.HandleFunc("/", auth)
	r.HandleFunc("/getUniversity", getUniversity)
	handler := c.Handler(r)
	fmt.Println("Server is listening...")
	log.Fatal(http.ListenAndServe(":3000", handler))
}

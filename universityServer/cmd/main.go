package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	usecase "universityServer/internal/tools/handle"

	"github.com/rs/cors"
)

func registration(w http.ResponseWriter, r *http.Request) {
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

func getUniversity(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		dataMap := make(map[string]int)

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

		jsonUniversity := make(map[string]string)
		jsonUniversity["name"] = result[0]
		jsonUniversity["description"] = result[1]
		jsonUniversity["imagePath"] = result[2]
		jsonUniversity["left"] = result[3]
		convertUniversity, err := json.Marshal(jsonUniversity)

		if err != nil {
			fmt.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
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

	r.HandleFunc("/signup", registration)
	r.HandleFunc("/getUniversity", getUniversity)
	handler := c.Handler(r)
	fmt.Println("Server is listening...")
	log.Fatal(http.ListenAndServe(":3000", handler))
}

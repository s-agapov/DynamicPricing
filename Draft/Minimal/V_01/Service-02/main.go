package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Data struct {
	Column1 string `json:"column_1" validate:"required"`
	Column2 int    `json:"column_2" validate:"min=0,max=100"`
	Column3 bool   `json:"column_3" validate:"boolean"`
}

func validateData(data Data) error {
	validate := validator.New()
	err := validate.Struct(data)
	if err != nil {
		return err
	}
	return nil
}

func validateHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	var data []Data
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid JSON payload", 400)
		return
	}

	for _, d := range data {
		err = validateData(d)
		if err != nil {
			log.Println(err)
			http.Error(w, "Invalid data payload", 400)
			return
		}
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func main() {
	http.HandleFunc("/validate", validateHandler)
	fmt.Println("Starting server at port 8082")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatal(err)
	}
}

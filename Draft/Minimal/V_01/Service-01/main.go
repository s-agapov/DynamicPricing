package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Data struct {
	Column1 string `json:"column_1"`
	Column2 string `json:"column_2"`
	Column3 string `json:"column_3"`
}

func readCSVFile(filePath string) ([]Data, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data := []Data{}

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return data, err
		}

		data = append(data, Data{
			Column1: record[0],
			Column2: record[1],
			Column3: record[2],
		})
	}

	return data, nil
}

func getDataHandler(w http.ResponseWriter, r *http.Request) {
	data, err := readCSVFile("data.csv")
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
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
	http.HandleFunc("/data", getDataHandler)
	fmt.Println("Starting service 01 at port 8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}

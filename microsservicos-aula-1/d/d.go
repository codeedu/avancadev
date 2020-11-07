package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":9093", nil)
}

type Result struct {
	Status string
}

func home(w http.ResponseWriter, r *http.Request) {

	log.Printf("Requisição recebida!")

	result := Result{Status: "teste microserviço 4"}

	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Error converting json")
	}

	fmt.Fprintf(w, string(jsonResult))

}

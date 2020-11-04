package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Cvv struct {
	Code string
}

func (c Cvv) Check() string {
	if c.Code == "123" {
		return "valid"
	}
	return "wrong"
}

type Result struct {
	Status string
}

func main() {

	http.HandleFunc("/", home)
	http.ListenAndServe(":9093", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	cvv := Cvv{
		Code: r.PostFormValue("cvvNumber"),
	}

	valid := cvv.Check()

	result := Result{Status: valid}

	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Error converting json")
	}

	fmt.Fprintf(w, string(jsonResult))

}

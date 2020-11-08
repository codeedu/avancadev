package main

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/go-retryablehttp"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Coupon struct {
	Code string
}

type Coupons struct {
	Coupon []Coupon
}

func (c Coupons) Check(code string) string {
	for _, item := range c.Coupon {
		if code == item.Code {
			return "valid"
		}
	}
	return "invalid"
}

type Result struct {
	Status string
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":9092", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	coupon := r.PostFormValue("coupon")
	result := makeHttpCall("http://localhost:9093", coupon)

	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Error converting json")
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(jsonResult))
}

type ValidResult struct {
	Valid bool `json:"valid"`
}

func makeHttpCall(urlMicroservice string, coupon string) Result {
	values := url.Values{}
	values.Add("coupon", coupon)

	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 5

	res, err := retryClient.PostForm(urlMicroservice, values)
	if err != nil {
		result := Result{Status: "Servidor fora do ar!"}
		return result
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error processing result")
	}

	response := ValidResult{}
	json.Unmarshal(data, &response)

	if response.Valid {
		return Result{ Status: "valid"}
	}

	return Result{ Status: "invalid"}
}

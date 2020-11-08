package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Valid bool `json:"valid"`
}

func (r Response) toString() string {
	jsonResult, err := json.Marshal(r)
	if err != nil {
		log.Fatal("Error converting json")
	}

	return string(jsonResult)
}

type Coupon struct {
	Code string
}

func main() {
	http.HandleFunc("/", process)
	http.ListenAndServe(":9093", nil)
}

var coupons = make(map[string]int)

func process(w http.ResponseWriter, r *http.Request) {
	res := Response{false}

	if r.PostFormValue("coupon") == "" {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, res.toString())
		return
	}

	coupon := Coupon{
		Code: r.PostFormValue("coupon"),
	}

	if coupons[coupon.Code] >= 3 {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, res.toString())
		return
	}

	coupons[coupon.Code] += 1

	log.Print(coupons)

	res = Response{true}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, res.toString())
}

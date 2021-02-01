package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Quote ... structure of api return
type Quote struct {
	Current       float64 `json:"c"`
	High          float64 `json:"h"`
	Low           float64 `json:"l"`
	Open          float64 `json:"o"`
	PreviousClose float64 `json:"pc"`
	Volume        int     `json:"t"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getStrings() (string, string) {
	return "&token=" + os.Getenv("FINN_KEY"),
		"https://finnhub.io/api/v1/quote?symbol="
}

func main() {
	var quote Quote
	var ticker = flag.String("t", "n/a", "Stock Ticker Lookup")
	key, url := getStrings()
	flag.Parse()
	r, err := http.Get(url + strings.ToUpper(*ticker) + key)
	check(err)

	body, err := ioutil.ReadAll(r.Body)
	check(err)
	err = json.Unmarshal(body, &quote)
	check(err)

	change := ((quote.Current - quote.PreviousClose) / quote.PreviousClose) * 100

	fmt.Printf("%s\nCurrent Price: $ %v \n", strings.ToUpper(*ticker), quote.Current)
	fmt.Printf("Previous Close: $ %v \n", quote.PreviousClose)

	if change >= 0 {
		fmt.Print(string("\033[32m"), "Percent Change: ")

	} else {
		fmt.Print(string("\033[31m"), "Percent Change: ")
	}
	fmt.Printf("%% %.3f \n", change)
	fmt.Print(string("\033[0m"), "\n")
}

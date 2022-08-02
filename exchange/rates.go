package exchange

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ExchangeRates struct {
	Base    string             `json:"base"`
	Date    string             `json:"date"`
	Rates   map[string]float64 `json:"rates"`
	Success bool               `json:"success"`
}

func RequestRatesFromAPI() ExchangeRates {
	url := "https://api.apilayer.com/exchangerates_data/latest?symbols=&base=IDR"

	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("apikey", "5dFvi1dREEw8zy6ZR4CkPoIahE8oT9kN")

	if err != nil {
		panic(err)
	}

	result, _ := client.Do(request)

	if result.Body != nil {
		defer result.Body.Close()
	}

	body, _ := ioutil.ReadAll(result.Body)

	var output ExchangeRates
	json.Unmarshal(body, &output)

	return output
}
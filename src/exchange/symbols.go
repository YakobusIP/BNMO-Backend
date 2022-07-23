package exchange

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type SymbolStruct struct {
	Success bool `json:"success"`
	Symbols interface{} `json:"symbols"`
}

func GetSymbolsFromAPI() SymbolStruct {
	url := "https://api.apilayer.com/exchangerates_data/symbols"

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

	var output SymbolStruct
	json.Unmarshal(body, &output)

	return output
}

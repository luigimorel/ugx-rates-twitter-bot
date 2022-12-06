package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

type CurrencyConvert struct {
	Result float64 `json:"result"`
}

func ConvertMoney() string {
	apiKey := os.Getenv("MONEY_API_KEY")

	url := "https://api.apilayer.com/currency_data/convert?to=UGX&from=USD&amount=1"

	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("apikey", apiKey)

	if err != nil {
		fmt.Println(err)
	}

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var result CurrencyConvert
	if err := json.Unmarshal(body, &result); err != nil {
		panic(err)
	}
	convertResultToInt := int(result.Result)

	convertResultToString := strconv.Itoa(convertResultToInt)

	currentDate := time.Now()
	formatDate := currentDate.Format("02 January 2006, 15:04:05 PM") + "\n"

	version := LogVersion()

	finalString := formatDate + "\n💰 1 USD ............ UGX " + convertResultToString + "\n\n" + version

	return fmt.Sprintf("%v", finalString)
}

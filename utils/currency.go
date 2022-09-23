package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-co-op/gocron"
)

type CurrencyConvert struct {
	Result float64 `json:"result"`
}

func RunCronJobs() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(1).Day().Do(func() {
		convertMoney()
	})

	s.StartBlocking()
}

func convertMoney() string {
	apiKey := os.Getenv("MONEY_API_KEY")

	url := "https://api.apilayer.com/fixer/convert?to=UGX&from=USD&amount=1"

	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("apikey", string(apiKey))

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
	convertResultToString := "\n💰 1 USD ............ UGX " + strconv.Itoa(convertResultToInt) + "\n\n"

	currentDate := time.Now()
	formatDate := currentDate.Format("Mon September, 15:04:05 PM") + "\n"
	version := LogVersion()

	finalString := formatDate + convertResultToString + version

	return fmt.Sprintf("%v", finalString)

}

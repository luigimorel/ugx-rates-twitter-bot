package client

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
	"github.com/morelmiles/ugx_rates/utils"
)

func Config() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		fmt.Println("Error loading .env file")
	}

	fmt.Println("Starting Server")
	scheduler := gocron.NewScheduler(time.UTC)

	apiKey := os.Getenv("API_KEY")
	apiKeySecret := os.Getenv("API_KEY_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")

	config := oauth1.NewConfig(apiKey, apiKeySecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	user, _, err := client.Accounts.VerifyCredentials(&twitter.AccountVerifyParams{})
	if err != nil {
		fmt.Printf("err : %v\n", err)
	}

	fmt.Printf("Account: @%s \n", user.ScreenName)
	scheduler.Every(1).Day().At("11:00").Do(func() {
		_, _, err = client.Statuses.Update(utils.ConvertMoney(), nil)
	})

	if err != nil {
		fmt.Printf("err : %v\n", err)
	}
}

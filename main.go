package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	consumerKey := os.Getenv("CONSUMERKEY")
	consumerSecret := os.Getenv("CONSUMERSECRET")
	accessToken := os.Getenv("ACCESSTOKEN")
	accessTokenSecret := os.Getenv("ACCESSTOKENSECRET")

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	user, _, err := client.Accounts.VerifyCredentials(&twitter.AccountVerifyParams{})
	if err != nil {
		fmt.Printf("Error connecting %v", err)
	}
	fmt.Printf("Account: %s (%s)\n", user.ScreenName, user.Name)

	messages, _, err := client.DirectMessages.EventsList(&twitter.DirectMessageEventsListParams{})
	if err != nil {
		fmt.Println("Error getting msgs")
	}
	for _, event := range messages.Events {
		fmt.Printf("%+v\n", event)
		fmt.Printf("  %+v\n", event.Message)
		fmt.Printf("  %+v\n", event.Message.Data)
	}

	event, _, err := client.DirectMessages.EventsShow("1066903366071017476", nil)
	fmt.Printf("DM Events Show:\n%+v, %v\n", event.Message.Data, err)
}

package settings

import (
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

type Client struct {
	Client *twitter.Client
}

func New() Client {
	c := Client{
		Client: connect(),
	}

	return c
}

func connect() *twitter.Client {
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

	return client
}

func (c *Client) GetCredentials() (string, error) {
	user, _, err := c.Client.Accounts.VerifyCredentials(&twitter.AccountVerifyParams{})
	if err != nil {
		return "", err
	}

	return user.ScreenName, nil
}

func (c *Client) GetMessages() ([]twitter.DirectMessageEvent, error) {
	messages, _, err := c.Client.DirectMessages.EventsList(&twitter.DirectMessageEventsListParams{})
	return messages.Events, err
}

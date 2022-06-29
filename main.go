package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/octavio-luna/twitter_bot/pkg/settings"
)

func main() {
	c := settings.New()

	// user, err := c.GetCredentials()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(user)

	// msgs, err := c.GetMessages()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for _, msg := range msgs {
	// 	fmt.Println(msg.Message.Data.Text)
	// }
	lambda.Start(c.GetCredentials)
}

/*
	user, _, err := client.Accounts.VerifyCredentials(&twitter.AccountVerifyParams{})
	if err != nil {
		fmt.Printf("Error connecting %v", err)
	}
	fmt.Printf("Account: %s (%s)\n", user.ScreenName, user.Name)

	messages, _, err := client.DirectMessages.EventsList(&twitter.DirectMessageEventsListParams{})
	if err != nil {
		return fmt.Sprint("Error getting msgs"), nil
	}
	// for _, event := range messages.Events {
	// 	fmt.Printf("%+v\n", event)
	// 	fmt.Printf("  %+v\n", event.Message)
	// 	fmt.Printf("  %+v\n", event.Message.Data.Text)
	// }
	return messages.Events[0].Message.Data.Text, nil
*/

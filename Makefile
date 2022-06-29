build: twitter_bot

twitter_bot: 
	GOOS=linux GOARCH=amd64 go build -o ./build/main
	zip ./build/main.zip ./build/main ./build/.env
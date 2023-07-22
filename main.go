package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	botToken := os.Getenv("YOUR_BOT_TOKEN")
	chatID := os.Getenv("YOUR_CHAT_ID")

	baseURL := "https://api.telegram.org/bot"
	endpoint := baseURL + botToken + "/sendMessage"

	message := "Hello, World!"
	resp, err := http.PostForm(
		endpoint,
		url.Values{
			"chat_id": {chatID},
			"text":    {message},
		},
	)

	if err != nil {
		fmt.Println("Error sending message:", err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Unexpected status code:", resp.StatusCode)
	}
	fmt.Println(resp.Body)
}

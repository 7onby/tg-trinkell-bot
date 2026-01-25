package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func sendTelegramMessage(text, botToken, chatID string) error {
	requestURL := "https://api.telegram.org/bot" + botToken + "/sendMessage"

	// Create the JSON payload
	values := map[string]string{"chat_id": chatID, "text": text}
	jsonParameters, err := json.Marshal(values)
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %w", err)
	}

	// Send the POST request
	res, err := http.Post(requestURL, "application/json", bytes.NewBuffer(jsonParameters))
	if err != nil {
		return fmt.Errorf("error sending HTTP request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %s", res.Status)
	}

	fmt.Println("Message sent successfully! Status:", res.Status)
	return nil
}

func main() {
	// Replace with your actual Bot Token and Chat ID
	botToken := ""
	chatID := "" //e.g., "123456789" or "-1001234567890" for channels
	messageText := "you are stupit!"

	if err := sendTelegramMessage(messageText, botToken, chatID); err != nil {
		log.Fatal(err)
	}
}

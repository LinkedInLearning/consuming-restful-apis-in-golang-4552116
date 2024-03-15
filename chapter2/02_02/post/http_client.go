package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	client := http.Client{
		Timeout: 3 * time.Second,
	}

	// Specify the URL for the POST/PUT request
	url := "https://jsonplaceholder.typicode.com/todos"

	// Define the Post request payload (data to be sent)
	payload := []byte(`{"title": "foo", "body": "bar", "userId": 1000}`)

	// Make the Post request
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error making Post request:", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error performing Post request:", err)
		return
	}

	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Print the response body
	fmt.Println("Post Response:", string(body))
}

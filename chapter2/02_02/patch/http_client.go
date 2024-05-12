package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	client := http.Client{
		Timeout: 3 * time.Second,
	}
	// Specify the URL for the Patch request
	url := "https://jsonplaceholder.typicode.com/posts/1"

	// Define the Patch request payload (data to be sent)
	payload := []byte(`{"userId": 1, "id": 1, "title": "sunt aut facere", "body": "quia et suscipit"}`)

	// Make the Patch request
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPatch, url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating Patch request:", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error performing Patch request:", err)
		return
	}

	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Print the response body
	fmt.Println("Patch Response:", string(body))
}

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
	// Specify the URL for the PUT request
	url := "https://jsonplaceholder.typicode.com/posts/1"

	// Define the Put request payload (data to be sent)
	payload := []byte(`{"userId": 1, "id": 1, "title": "sunt aut facere", "body": "quia et suscipit"}`)

	// Make the Put request
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPut, url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating Put request:", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error performing Put request:", err)
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
	fmt.Println("Put Response:", string(body))
}

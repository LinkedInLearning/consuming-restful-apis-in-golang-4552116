package main

import (
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
	// Specify the URL for the Delete request
	url := "https://jsonplaceholder.typicode.com/posts/1"

	// Make the Delete request
	req, err := http.NewRequestWithContext(context.Background(), http.MethodDelete, url, nil)
	if err != nil {
		fmt.Println("Error creating Delete request:", err)
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

	fmt.Println("Response Code: ", resp.StatusCode)

	// Print the response body
	fmt.Println("Delete Response:", string(body))
}

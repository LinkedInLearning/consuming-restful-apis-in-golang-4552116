package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type commentsResponse struct {
	PostId int    `json:"postId"`
	ID     int    `json:"id"`
	Email  int    `json:"email"`
	Name   int    `json:"name"`
	Body   string `json:"body"`
}

type app struct {
	client http.Client
}

func (a app) performRequest(req *http.Request) (map[string]interface{}, error) {
	resp, err := a.client.Do(req)
	if err != nil {
		fmt.Println("Could not complete request due to err: ", err)
		// Wrap the original error with additional context
		return nil, fmt.Errorf("encountered an error while performing the request: %w", err)
	}

	defer resp.Body.Close()

	var comment map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&comment)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return comment, err
	}

	return comment, nil
}

func main() {

	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://jsonplaceholder.typicode.com/comments/1", nil)
	if err != nil {
		fmt.Println("Could not create request due to err: ", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Could not complete request due to err: ", err)
		return
	}
	defer resp.Body.Close()

	// Decode the JSON response into a map
	var comment map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&comment)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Printf("Response Status Code: %d\n", resp.StatusCode)

	// Print the response
	fmt.Printf("Response Body: %+v \n", comment)

}

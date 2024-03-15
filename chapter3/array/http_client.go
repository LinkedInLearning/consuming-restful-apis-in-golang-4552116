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
	Email  string `json:"email"`
	Name   string `json:"name"`
	Body   string `json:"body"`
}

func main() {

	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://jsonplaceholder.typicode.com/comments", nil)
	if err != nil {
		fmt.Println("Could not create the request due to: ", err.Error())
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Could not perform the request due to: ", err.Error())
	}

	defer resp.Body.Close()

	// Decode the JSON response
	var comments []commentsResponse
	err = json.NewDecoder(resp.Body).Decode(&comments)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Printf("Response Status Code: %d\n", resp.StatusCode)

	// Print the response
	fmt.Printf("Number of comments: %+v \n", len(comments))
	fmt.Printf("Comments[0]: %+v \n", comments[0])

}

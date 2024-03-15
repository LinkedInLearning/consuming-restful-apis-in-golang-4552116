package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://jsonplaceholder.typicode.com/todos/1", nil)
	if err != nil {
		fmt.Println("Error occured while creating request: ", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error occured while performing request: ", err)
		return
	}

	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Printf("Response Status Code: %d\n", resp.StatusCode)

	// Print the response
	fmt.Println("Response Body:", string(body))

}

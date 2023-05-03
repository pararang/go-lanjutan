package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Create an HTTP client
	client := &http.Client{}

	// Send a GET request to the /ping endpoint on the server
	resp, err := client.Get("http://localhost:8081/hi/mid")
	// resp, err := client.Get("http://localhost:8081/ping")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the response body as a string
	fmt.Println(string(body))
}

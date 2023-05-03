package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Player struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// Read the contents of player.json into a byte array
	fileBytes, err := ioutil.ReadFile("player.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Unmarshal the JSON data into a slice of MyData structs
	var data []Player
	err = json.Unmarshal(fileBytes, &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Print the data for each object
	for _, d := range data {
		fmt.Println("Name:", d.Name)
		fmt.Println("Email:", d.Email)
		fmt.Println("---------------")
	}

	FreeText()
}

func FreeText(){
	content, err := ioutil.ReadFile("puisi.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("File contents:")
	fmt.Println(string(content))
}

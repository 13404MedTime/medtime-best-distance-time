package main

import (
	"fmt"
	"handler/function"
	"io"
	"os"
)

// this part is to test the function locally
func main() {
	// Open and read the JSON file
	jsonFile, err := os.Open("request.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// Call the handler
	handler := function.Handle(byteValue)

	// Print the response
	fmt.Printf("Body: %s\n", handler)
}

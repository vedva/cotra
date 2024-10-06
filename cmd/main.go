package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Set up the route and handler
	http.HandleFunc("/service", serviceHandler)

	// Start the server on port 8080
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

// Function to calculate the sum of ASCII values of the characters in the name
func calculateASCIISum(name string) int {
	sum := 0
	for _, char := range name {
		sum += int(char)
	}
	return sum
}

// Handler function for GET requests
func serviceHandler(w http.ResponseWriter, r *http.Request) {
	// Get the 'name' parameter from the query
	fmt.Println("Servicing request")
	name := r.URL.Query().Get("name")

	if name == "" {
		http.Error(w, "Missing 'name' parameter", http.StatusBadRequest)
		return
	}

	// Calculate the ASCII sum
	asciiSum := calculateASCIISum(name)

	// Create the response message
	response := fmt.Sprintf("Hello %s, ASCII sum: %d", name, asciiSum)

	// Write the response
	w.Write([]byte(response))
}

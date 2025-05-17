package main

import (
	"fmt"
	"net/http"
)

func main() {

	// HTTP WRITER -> to handle HTTP response
	// HTTP REQUEST -> to handle HTTP request
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	// Listen and serve on port 8080
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}

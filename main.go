package main

import (
	"fmt"
	"github.com/hussamElmaghraby/store_hub_go/internal/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.InitRouter()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	// Listen and serve on port 8080
	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

	//http.ListenAndServe(":8080", nil)
}

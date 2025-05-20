package main

import (
	"fmt"
	"github.com/hussamElmaghraby/store_hub_go/configs"
	"github.com/hussamElmaghraby/store_hub_go/internal/database"
	"github.com/hussamElmaghraby/store_hub_go/internal/routes"
	"log"
	"net/http"
)

func main() {
	configs.LoadEnv()

	database.InitDB()

	port := configs.GetEnv("PORT", "8080")
	appName := configs.GetEnv("APP_NAME", "No Name")

	router := routes.InitRouter()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	// Listen and serve on port 8080
	fmt.Printf("%s Server started on //localhost:%s\n", appName, port)
	log.Fatal(http.ListenAndServe(":"+port, router))

	//http.ListenAndServe(":8080", nil)
}

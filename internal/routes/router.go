package routes

import (
	"github.com/gorilla/mux"
	"github.com/hussamElmaghraby/store_hub_go/internal/handlers"
)

func InitRouter() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	router.HandleFunc("/products", handlers.CreatProductHandler).Methods("POST")
	router.HandleFunc("/products", handlers.GetAllProducts).Methods("GET")
	router.HandleFunc("/products/{id}", handlers.GetProductById).Methods("GET")
	router.HandleFunc("/products/{id}", handlers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", handlers.DeleteProduct).Methods("DELETE")

	return router
}

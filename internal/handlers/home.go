package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/hussamElmaghraby/store_hub_go/internal/database"
	"github.com/hussamElmaghraby/store_hub_go/internal/models"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Handler - Home Page!")
}

func CreatProductHandler(w http.ResponseWriter, r *http.Request) {

	//Declares a variable with a type, zero-initialized
	var product models.Product

	//Declares and initializes a variable with a value
	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		http.Error(w, "Failed to create body", http.StatusInternalServerError)
		return
	}

	result := database.DB.Create(&product)
	if result.Error != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	result := database.DB.Find(&products)
	if result.Error != nil {
		http.Error(w, "Failed to get products", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

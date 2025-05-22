package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/hussamElmaghraby/store_hub_go/internal/database"
	"github.com/hussamElmaghraby/store_hub_go/internal/models"
	"net/http"
	"strconv"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	// extract the id from the url parameters
	idStr := vars["id"]

	// covert ID to int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	// find the product by ID
	var product models.Product
	result := database.DB.First(&product, id)
	if result.Error != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	// decode the request body into the product struct
	// and update the product
	var updatedProduct models.Product
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	// Basic validation:  name must not be empty
	if updatedProduct.Name == "" || updatedProduct.Price == 0 || updatedProduct.Quantity == 0 {
		http.Error(w, "Name, Price and Quantity must not be empty", http.StatusBadRequest)
		return
	}

	// update the product
	result = database.DB.Model(&product).Updates(updatedProduct)
	if result.Error != nil {
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		return
	}
	// return the updated product
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)

}

package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/NickSFU/ProductsAPI/db"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	products := db.GetProducts()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	idStr := strings.TrimPrefix(path, "/product/")
	if idStr == "" {
		products := db.GetProducts()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}
	product := db.GetProductByID(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

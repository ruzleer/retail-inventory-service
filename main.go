package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"


)

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var products = []Product{
	{ID: 1, Name: "Дрель"},
	{ID: 2, Name: "Отвертка"},
	{ID: 3, Name: "Молоток"},
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func productHandler(w http.ResponseWriter, r *http.Request) {

	url := r.URL.Path
	content := strings.Split(url, "/")
	id, err := strconv.Atoi(content[2])

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, product := range products {
		if product.ID == id {
			json.NewEncoder(w).Encode(product)
			return
		}
	}

//	w.WriteHeader(http.StatusNotFound)
	http.Error(w, "{\"error\" : \"Продукт не найден\"}", http.StatusNotFound)
//	w.Write([]byte(`{"error": "Продукт не найден"}`))

}

func createProductHandler(w http.ResponseWriter, r *http.Request) {

	var newProduct Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)

	if( err != nil) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if( newProduct.Name == "") {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "{\"error\" : \"Имя продукта не может быть пустым\"}", http.StatusBadRequest)
		return
	}

	newProduct.ID = len(products) + 1
	products = append(products, newProduct)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newProduct)

}

func main() {
	http.HandleFunc("GET /health", healthHandler)
	http.HandleFunc("GET /products", productsHandler)
	http.HandleFunc("POST /products", createProductHandler)
	http.HandleFunc("GET /product/{id}", productHandler)

	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}

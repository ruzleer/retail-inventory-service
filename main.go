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

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func productsHandler(w http.ResponseWriter, r *http.Request) {

	products := []Product{
		{ID: 1, Name: "Дрель"},
		{ID: 2, Name: "Отвертка"},
		{ID: 3, Name: "Молоток"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func productHandler(w http.ResponseWriter, r *http.Request) {

	products := []Product{
		{ID: 1, Name: "Дрель"},
		{ID: 2, Name: "Отвертка"},
		{ID: 3, Name: "Молоток"},
	}

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
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/products", productsHandler)
	http.HandleFunc("/product/{id}", productHandler)

	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}

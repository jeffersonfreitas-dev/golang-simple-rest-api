package handlers

import (
	"encoding/json"
	"net/http"
	"simple-rest-api/internal/models"

	"github.com/google/uuid"
)

func (h Handlers) registerProductEndpoints() {
	http.HandleFunc("GET /products", h.getAllProducts)
	http.HandleFunc("GET /products/{id}", h.getProductById)
	http.HandleFunc("POST /products", h.addProduct)
}

func (h Handlers) getAllProducts(w http.ResponseWriter, r *http.Request) {
	products := h.productCases.GetAll()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func (h Handlers) addProduct(w http.ResponseWriter, r *http.Request) {
	var body models.ProductRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	response := h.productCases.Add(body)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h Handlers) getProductById(w http.ResponseWriter, r *http.Request) {
	idRequest := r.PathValue("id")

	id, err := uuid.Parse(idRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	product, err := h.productCases.FindById(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

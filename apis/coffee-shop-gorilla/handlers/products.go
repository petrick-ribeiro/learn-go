package handlers

import (
	"coffee-shop/data"
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// GET function returns the products from the data store.
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// fetch
	lp := data.GetProducts()

	// serialize the list to JSON
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}

// POST function add a new product.
func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

  prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}

// PUT function update specified product.
func (p Products) UpdateProducts(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id, err := strconv.Atoi(vars["id"])
  if err != nil {
    http.Error(w, "Unable to convert id", http.StatusBadRequest)
    return
  }

	p.l.Println("Handle PUT Product", id)
  prod := r.Context().Value(KeyProduct{}).(data.Product)

  err = data.UpdateProduct(id, &prod)
  if err == data.ErrProductNotFound {
    http.Error(w, "Product not found", http.StatusNotFound)
    return
  }

  if err != nil {
    http.Error(w, "Product not found", http.StatusInternalServerError)
    return
  }
}

type KeyProduct struct{}

// Middleware
func (p Products) ValidateProduct(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    prod := data.Product{}

    err := prod.FromJSON(r.Body)
    if err != nil {
      http.Error(w, "Error reading the product", http.StatusBadRequest)
      return
    }

    ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
    req := r.WithContext(ctx) 
    
    // Call the next handler.
    next.ServeHTTP(w, req)
  })
}

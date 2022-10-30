package handlers

import (
	"coffee-shop/data"
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

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
	}

	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
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

	prod := &data.Product{}

	err = prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
	}

  err = data.UpdateProduct(id, prod)
  if err == data.ErrProductNotFound {
    http.Error(w, "Product not found", http.StatusNotFound)
    return
  }

  if err != nil {
    http.Error(w, "Product not found", http.StatusInternalServerError)
    return
  }
}

package handlers

import (
	"coffee-shop/data"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  // handle the request for a list of products
  // GET
  if r.Method == http.MethodGet {
    p.getProducts(w, r)
    return
  }

  // POST
  if r.Method == http.MethodPost {
    p.addProduct(w, r)
    return
  }

  // PUT
  if r.Method == http.MethodPut {
    // 
  }

  // catch all
  // if no method is satisfied return an error
  w.WriteHeader(http.StatusMethodNotAllowed)
}

// GET function returns the products from the data store
func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
  p.l.Println("Handle GET Products")

  // fetch
  lp := data.GetProducts()

  // serialize the list to JSON
  err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
  p.l.Println("Handle POST Product")

  prod := &data.Product{}  
  err := prod.FromJSON(r.Body)
  if err != nil {
    http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
  }

  p.l.Printf("Prod: %#v", prod)
  data.AddProduct(prod)
}

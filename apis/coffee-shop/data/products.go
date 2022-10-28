package data

import (
	"encoding/json"
	"io"
	"time"
)

// Create the structure of data for API
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdateOn    string  `json:"-"`
	DeleteOn    string  `json:"-"`
}

func (p *Product) FromJSON(r io.Reader) error {
  e := json.NewDecoder(r)
  return e.Decode(p)
}

// Collection of Product
type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
  e := json.NewEncoder(w)
  return e.Encode(p)
}

func GetProducts() Products {
  return productList
}

func AddProduct(p *Product)  {
  p.ID = getNextID()
  productList = append(productList, p)
}

func getNextID() int {
  lp := productList[len(productList) -1]
  return lp.ID + 1
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffe",
		Price:       2.34,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffe without milk",
		Price:       1.34,
		SKU:         "efg321",
		CreatedOn:   time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	},
}

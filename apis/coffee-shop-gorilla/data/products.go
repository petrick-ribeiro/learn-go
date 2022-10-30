package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

// Create the structure of data for API
type Product struct {
	ID          int     `json:"id"`
  Name        string  `json:"name"  validate:"required"`
	Description string  `json:"description"`
  Price       float32 `json:"price" validate:"gt=0"`
  SKU         string  `json:"sku"   validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdateOn    string  `json:"-"`
	DeleteOn    string  `json:"-"`
}

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *Product) Validate() error {
  validate := validator.New()
  validate.RegisterValidation("sku", validateSKU)

  return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
  // SKU format is abc-def-123
  reg     := regexp.MustCompile(`[a-z]+-[a-z]+-[0-9]+`)
  matches := reg.FindAllString(fl.Field().String(), -1)

  if len(matches) != 1 {
    return false
  }

  return true
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

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
  _, pos, err := findProduct(id)
  if err != nil {
    return err
  }

  p.ID = id
  productList[pos] = p

  return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error)  {
  for i, p := range productList {
    if p.ID == id {
      return p, i, nil
    }
  }
  return nil, -1, ErrProductNotFound
}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

// Data Source
var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.34,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.34,
		SKU:         "efg321",
		CreatedOn:   time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	},
}

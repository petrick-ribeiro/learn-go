package data

import "testing"

func TestValidation(t *testing.T)  {
  p := &Product{
    Name: "Coffee",
    Price: 1.5,
    SKU: "abc-cde-123",
  }

  err := p.Validate()

  if err != nil {
    t.Fatal(err)
  }
}

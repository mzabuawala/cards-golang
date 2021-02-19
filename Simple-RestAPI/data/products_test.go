package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		ID:          1,
		Name:        "Test",
		Description: "test description",
		Price:       1.5,
		SKU:         "test-123",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}

package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

// Product defines the structure for an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float32 `json:"price" validate:"required,gt=1"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// FromJSON will convert it to JSON
func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

// Validate will validate the product
func (p *Product) Validate() error {
	validate := validator.New()
	// Here we are registering external validator function
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

// validateSKU it will validate SKU
func validateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[a-z]+-[0-9]+`)
	matches := re.FindAllString((fl.Field().String()), -1)
	if len(matches) != 1 {
		return false
	}
	return true
}

// Products is a collection of Product
type Products []*Product

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
//
// https://golang.org/pkg/encoding/json/#NewEncoder
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// GetProducts returns a list of products
func GetProducts() Products {
	return productList
}

// AddProduct will add a new product
func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

// UpdateProduct will update existing product
func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p

	return nil
}

// RemoveProduct will remove existing product
func RemoveProduct(id int) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	productList = append(productList[:pos], productList[pos+1:]...)

	return nil
}

// ErrProductNotFound is a format error string
var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
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

// productList is a hard coded list of products for this
// example data source
var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc-323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd-34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

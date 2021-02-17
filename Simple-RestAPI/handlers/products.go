package handlers

import (
	"context"
	"log"
	"net/http"
	"product-api/data"
	"strconv"

	"github.com/gorilla/mux"
)

// Products is a http.Handler
type Products struct {
	l *log.Logger
}

// NewProducts creates a products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// GetProducts returns the products from the data store
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// fetch the products from the datastore
	lp := data.GetProducts()

	// serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// AddProduct returns the products from the data store
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddProduct(&prod)
}

// UpdateProduct returns the products from the data store
func (p Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Product")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	err := data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

// RemoveProduct returns the products from the data store
func (p Products) RemoveProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle DELETE Product")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	err := data.RemoveProduct(id)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

// KeyProduct is a key struct for request context
type KeyProduct struct{}

// MiddlewareValidateProduct Middleware for out API
func (p Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(rw http.ResponseWriter, r *http.Request) {
			p.l.Println("In the product validation middleware")
			prod := data.Product{}
			err := prod.FromJSON(r.Body)
			if err != nil {
				p.l.Println("[ERROR]: desearlizing product", err)
				http.Error(rw, "Error reading product", http.StatusBadRequest)
				return
			}
			ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
			r = r.WithContext(ctx)
			next.ServeHTTP(rw, r)
		})
}

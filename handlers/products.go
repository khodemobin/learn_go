package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/khodemobin/learn_go/data"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
}

func (products *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	products.logger.Println("Handle GET Products")

	// fetch the products from the datastore
	lp := data.GetProducts()

	// serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (products *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	products.logger.Println("Handle Post Products")

	prod := &data.Product{}
	err := prod.FromJson(r.Body)

	if err != nil {
		http.Error(w, "Unable to parse json", http.StatusBadRequest)

	}

	products.logger.Printf("Prob: %#v", prod)

	data.AddProduct(prod)

}

func (p Products) UpdateProducts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to parse id", http.StatusBadRequest)
		return
	}

	p.logger.Println("Handle update product")

	prod := &data.Product{}

	err = prod.FromJson(r.Body)

	if err != nil {
		http.Error(w, "Unable to parse json", http.StatusBadRequest)
		return
	}

	err = data.UpdateProduct(id, prod)

	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

}

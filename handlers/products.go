package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/khodemobin/learn_go/data"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
}

func (products *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		products.getProducts(w, r)
		return
	}

	if r.Method == http.MethodPost {
		products.addProduct(w, r)
		return
	}

	if r.Method == http.MethodPut {

		regex := regexp.MustCompile(`/([0-9]+)`)
		g := regex.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 1 {
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, _ := strconv.Atoi(idString)

		products.logger.Println("got id", id)
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

}

func (products *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	products.logger.Println("Handle GET Products")

	// fetch the products from the datastore
	lp := data.GetProducts()

	// serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (products *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	products.logger.Println("Handle Post Products")

	prod := &data.Product{}
	err := prod.FromJson(r.Body)

	if err != nil {
		http.Error(w, "Unable to parse json", http.StatusBadRequest)

	}

	products.logger.Printf("Prob: %#v", prod)

	data.AddProduct(prod)

}

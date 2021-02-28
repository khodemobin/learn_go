package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/khodemobin/learn_go/handlers"
)

func ProductRoutes(sm *mux.Router, logger *log.Logger) {
	productsHandler := handlers.NewProducts(logger)

	sm.HandleFunc("/", productsHandler.GetProducts).Methods(http.MethodGet)
	sm.HandleFunc("/{id:[0-9]+}", productsHandler.UpdateProducts).Methods(http.MethodPut)
	sm.HandleFunc("/", productsHandler.AddProduct).Methods(http.MethodPost)
}

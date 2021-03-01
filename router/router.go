package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/khodemobin/learn_go/handlers"
)

func RegisterRoutes(sm *mux.Router, logger *log.Logger) {
	productsHandler := handlers.NewProducts(logger)
	sm.HandleFunc("/", productsHandler.GetProducts).Methods(http.MethodGet)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", productsHandler.UpdateProducts)
	putRouter.Use(productsHandler.MiddlewareValidateProduct)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", productsHandler.AddProduct)
	postRouter.Use(productsHandler.MiddlewareValidateProduct)

}

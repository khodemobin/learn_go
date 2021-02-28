package router

import (
	"log"

	"github.com/gorilla/mux"
)

func RegisterRoutes(sm *mux.Router, logger *log.Logger) {

	ProductRoutes(sm, logger)

}
